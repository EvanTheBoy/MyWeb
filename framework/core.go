package framework

import (
	"log"
	"net/http"
	"strings"
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

// Get 匹配GET方法, 增加路由规则
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// Post 匹配POST方法, 增加路由规则
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// Put 匹配PUT方法, 增加路由规则
func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// Delete 匹配DELETE方法, 增加路由规则
func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// FindRouterByRequest 根据请求查找路由
func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
	method := request.Method // 获取请求方法
	url := request.URL.Path
	upperMethod := strings.ToUpper(method)
	if methodHandler, ok := c.router[upperMethod]; ok {
		return methodHandler.FindHandler(url)
	}
	return nil
}

// 负责路由分发, 所有请求都进入这个函数
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
