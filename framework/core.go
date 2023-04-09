package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree    // 所有的路由
	middlewares []ControllerHandler // 从core这边设置的中间件
}

func NewCore() *Core {
	// 不使用哈希表来进行二级映射了,因为需要使用动态路由匹配, 之前那种方法不奏效了
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{
		router: router,
	}
}

// RegisterMiddleware 注册中间件
func (c *Core) RegisterMiddleware(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

// Get 匹配GET方法, 增加路由规则
func (c *Core) Get(uri string, handlers ...ControllerHandler) {
	// 将core的middlewares和handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// Post 匹配POST方法, 增加路由规则
func (c *Core) Post(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// Put 匹配PUT方法, 增加路由规则
func (c *Core) Put(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// Delete 匹配DELETE方法, 增加路由规则
func (c *Core) Delete(uri string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(uri, allHandlers); err != nil {
		log.Fatal("Add router error:", err)
	}
}

// FindNodeByRequest 根据请求查找路由
// 我们修改了这个函数的返回值和实现方式, 现在我们返回的是匹配的节点
func (c *Core) FindNodeByRequest(request *http.Request) *Node {
	method := request.Method // 获取请求方法
	url := request.URL.Path
	upperMethod := strings.ToUpper(method)
	if methodTree, ok := c.router[upperMethod]; ok {
		return methodTree.root.matchNode(url)
	}
	return nil
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 负责路由分发, 所有请求都进入这个函数
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 封装自定义的Context
	ctx := NewContext(request, response)
	// 寻找路由
	node := c.FindNodeByRequest(request)
	if node == nil {
		ctx.SetStatus(404).Json("not found")
		return
	}
	// 设置控制器
	ctx.SetHandlers(node.handlers)
	// 设置路由参数
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)
	// 调用路由函数, 若返回err则表示内部错误, 应该给一个500的状态码
	// 控制器的index是从0开始的, 所以这里调用Next()就是调用的当前的控制器
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		return
	}
}
