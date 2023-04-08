package framework

type Group struct {
	core        *Core
	parent      *Group              // 如果有的话, 指向上一个Group
	prefix      string              // 该group的通用前缀
	middlewares []ControllerHandler // 存放中间件
}

// NewGroup 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:        core,
		parent:      nil,
		prefix:      prefix,
		middlewares: []ControllerHandler{},
	}
}

// RegisterMiddleware 注册中间件
func (g *Group) RegisterMiddleware(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

/*
接下来实现IGroup接口
*/

// InitGroup 实现InitGroup方法
func (g *Group) InitGroup(url string) IGroup {
	cGroup := NewGroup(g.core, url)
	cGroup.parent = g
	return cGroup
}

// 获取当前Group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// 获取某个group的middleware
// 这里指的是获取除了Get/Post/Put/Delete之外设置的middlewares
func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.middlewares, g.middlewares...)
}

// Get 实现Get方法
func (g *Group) Get(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	allHandlers := append(g.middlewares, handlers...)
	g.core.Get(url, allHandlers...)
}

// Post 实现Post方法
func (g *Group) Post(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	allHandlers := append(g.middlewares, handlers...)
	g.core.Post(url, allHandlers...)
}

// Put 实现Put方法
func (g *Group) Put(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	allHandlers := append(g.middlewares, handlers...)
	g.core.Put(url, allHandlers...)
}

// Delete 实现Delete方法
func (g *Group) Delete(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	allHandlers := append(g.middlewares, handlers...)
	g.core.Delete(url, allHandlers...)
}
