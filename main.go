package main

import (
	"github.com/gohade/my-web/app/console"
	"github.com/gohade/my-web/app/http"
	"github.com/gohade/my-web/framework"
	"github.com/gohade/my-web/framework/provider/app"
	"github.com/gohade/my-web/framework/provider/kernel"
)

func main() {
	// 初始化容器
	container := framework.NewHadeContainer()
	// 绑定APP服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 将HTTP引擎初始化, 并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	// 运行root命令
	console.RunCommand(container)
}
