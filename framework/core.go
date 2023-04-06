package framework

import (
	"log"
	"net/http"
)

type Core struct {
	// 路由url对应的handler处理函数, 这里用一个map结构把它们绑定起来
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(request *http.Request, response http.ResponseWriter) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(request, response)

	// 一个简单的路由选择器, 暂时先这么写
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")
	err := router(ctx)
	if err != nil {
		return
	}
}
