package router

import (
	"MyWeb/app"
	"MyWeb/framework"
)

func RegisterRoute(c *framework.Core) {
	// 路由的url和对应的处理函数
	c.Get("foo", app.FooControllerHandler)
}
