package main

import (
	"context"
	"github.com/gohade/my-web/framework/gin"
	"github.com/gohade/my-web/framework/middleware"
	"github.com/gohade/my-web/provider/demo"
	"github.com/gohade/my-web/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	err1 := core.Bind(&demo.DemoServiceProvider{})
	if err1 != nil {
		return
	}
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())
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
	// 一般想要优雅关闭, 我们都会需要用到WithTimeout这个函数, 并给它设置一个阈值时间, 避免一直等待
	// Shutdown这个函数, 它并不会强制停止当前正在执行的进程, 它一定是要先等待, 等待现在正在执行任务的
	// 进程结束, 他才会真的结束. 所以我们在浏览器输入访问一个需要执行10s的地址, 并在控制台按ctrl+C键
	// 想要强制终止进程的时候, 我们会发现并没有用, 因为Shutdown需要等待当前正在执行的进程结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	if err2 := server.Shutdown(timeoutCtx); err2 != nil {
		log.Fatal("Server shutdown:", err2)
	}
}
