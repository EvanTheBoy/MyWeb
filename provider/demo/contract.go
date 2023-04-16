package demo

// Key Demo服务的key
// 定义服务的关键字凭证
const Key = "hade:demo"

// Foo Demo服务接口定义的一个数据结构
type Foo struct {
	Name string
}

// Service Demo服务的接口
type Service interface {
	GetFoo() Foo
}
