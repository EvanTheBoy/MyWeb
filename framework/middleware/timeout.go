package middleware

import (
	"MyWeb/framework"
	"context"
	"log"
	"time"
)

// Timeout 超时控制逻辑, 改写
func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 调用具体的逻辑
		err := c.Next()
		if err != nil {
			return err
		}

		select {
		case p := <-panicChan:
			c.SetStatus(500).Json("time out")
			log.Println(p)
		case <-finish:
			log.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.SetStatus(500).Json("time out")

		}
		return nil
	}
}
