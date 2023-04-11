package main

import (
	"MyWeb/framework"
	"MyWeb/framework/middleware"
	"MyWeb/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.RegisterMiddleware(middleware.Recovery())
	core.RegisterMiddleware(middleware.Cost())
	//core.RegisterMiddleware(middleware.Timeout(1 * time.Second))
	router.RegisterRoute(core)
	server := http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	// 将启动的服务的代码单独用一个goroutine去管理
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	// 在main函数所在的goroutine中监听信号
	quit := make(chan os.Signal)
	// 监控信号SIGINT, SIGTERM和SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 阻塞在这里
	<-quit

	// 调用server.shutdown()方法结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
}
