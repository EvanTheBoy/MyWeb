package framework

// NewInstance 定义一个新的实例
type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider 定义了服务提供者需要实现的接口
type ServiceProvider interface {
	Name() string                   // 代表服务提供者的凭证
	Register(Container) NewInstance // 在服务容器中注册一个实例化服务的方法
	Boot(Container) error           // 实例化服务, 包括一些基础配置, 初始化参数的操作
	IsDefer() bool                  // 决定是否在注册的时候实例化服务
	Params(Container) []interface{} // 定义传递给NewInstance的参数
}
