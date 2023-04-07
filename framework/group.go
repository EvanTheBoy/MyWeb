package framework

type Group struct {
	core   *Core
	parent *Group // 如果有的话, 指向上一个Group
	prefix string // 该group的通用前缀
}

// NewGroup 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// IGroup 为了提升扩展性, 届时我们在修改了Group的具体实现后, 只返回这个接口
type IGroup interface {
	// Get Post Put Delete 实现HTTPMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	InitGroup(string) IGroup
}

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

// Get 实现Get方法
func (g *Group) Get(url string, handler ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	g.core.Get(url, handler)
}

// Post 实现Post方法
func (g *Group) Post(url string, handler ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	g.core.Post(url, handler)
}

// Put 实现Put方法
func (g *Group) Put(url string, handler ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	g.core.Put(url, handler)
}

// Delete 实现Delete方法
func (g *Group) Delete(url string, handler ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	g.core.Delete(url, handler)
}
