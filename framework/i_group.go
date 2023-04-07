package framework

// IGroup 为了提升扩展性, 我们对外界主要提供返回这个接口
// 这样如果需要修改Group, 可以减少工作量, 同时也可以减少模块之间的耦合性
type IGroup interface {
	// Get Post Put Delete 实现HTTPMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	InitGroup(string) IGroup
}
