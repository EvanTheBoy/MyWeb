package app

import (
	"MyWeb/framework"
	"context"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	finish := make(chan struct{}, 1)       // 用来通知结束的管道
	panicChan := make(chan interface{}, 1) // 通知panic异常的管道
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), 1*time.Second)
	defer cancel()

	// 开启协程
	go func() {
		defer func() {
			if p := recover(); p != nil {
				// 能进入这个if语句的p都是异常的, 所以通通放进panicChan里面
				panicChan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		err := ctx.Json(200, "ok")
		if err != nil {
			return
		}
		// 操作完了后, 还要记得通知结束
		finish <- struct{}{}
	}()

	select {
	// 监听异常事件
	case p := <-panicChan:
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		log.Println(p)
		err1 := ctx.Json(500, "panic")
		if err1 != nil {
			return err1
		}
		// 监听结束事件
	case <-finish:
		fmt.Println("finish")
		// 监听超时事件
	case <-durationCtx.Done():
		ctx.WriteMux().Lock()
		defer ctx.WriteMux().Unlock()
		err2 := ctx.Json(500, "time out")
		if err2 != nil {
			return err2
		}
		ctx.SetHasTimeout()
	}
	return nil
}
