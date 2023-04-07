package framework

import (
	"log"
	"net/http"
)

type Core struct {
	// 这里是一个二级映射, 目的是匹配路由
	router map[string]*Tree
}

func NewCore() *Core {
	// 不使用哈希表来进行二级映射了,因为需要使用动态路由匹配, 之前那种方法不奏效了
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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
