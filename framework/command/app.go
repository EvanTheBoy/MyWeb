package command

import (
	"context"
	"github.com/gohade/my-web/framework/cobra"
	"github.com/gohade/my-web/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// initAppCommand 初始化app命令和其子命令
func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}

// AppCommand 是命令行参数第一级为app的命令，它没有实际功能，只是打印帮助文档
var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	Long:  "业务应用控制命令，其包含业务启动，关闭，重启，查询等功能",
	RunE: func(c *cobra.Command, args []string) error {
		// 打印帮助文档
		c.Help()
		return nil
	},
}

// appStartCommand 启动一个Web服务
var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个Web服务",
	RunE: func(c *cobra.Command, args []string) error {
		// 从Command中获取服务容器
		container := c.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()
		// 创建一个Server服务
		server := &http.Server{
			Handler: core,
			Addr:    ":8080",
		}
		// 这个goroutine是启动服务的goroutine
		go func() {
			server.ListenAndServe()
		}()
		// 当前的goroutine等待信号量
		quit := make(chan os.Signal)
		// 监控信号：SIGINT, SIGTERM, SIGQUIT
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		// 这里会阻塞当前goroutine等待信号
		<-quit
		// 调用Server.Shutdown graceful结束
		// 一般想要优雅关闭, 我们都会需要用到WithTimeout这个函数, 并给它设置一个阈值时间, 避免一直等待
		// Shutdown这个函数, 它并不会强制停止当前正在执行的进程, 它一定是要先等待, 等待现在正在执行任务的
		// 进程结束, 他才会真的结束. 所以我们在浏览器输入访问一个需要执行10s的地址, 并在控制台按ctrl+C键
		// 想要强制终止进程的时候, 我们会发现并没有用, 因为Shutdown需要等待当前正在执行的进程结束
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		return nil
	},
}
