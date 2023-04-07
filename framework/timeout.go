package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

// TimeoutHandler 超时的中间件
func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	// 这里使用的是回调函数
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(durationCtx)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的业务逻辑
			err := fun(c)
			if err != nil {
				return
			}

			finish <- struct{}{}
		}()
		// 执行业务逻辑后的操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			_, err := c.responseWriter.Write([]byte("Time out"))
			if err != nil {
				return err
			}
		}
		return nil
	}
}
