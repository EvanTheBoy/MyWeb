package framework

import "sync"

type Container interface {
	Bind(provider ServiceProvider) error                           // 注册方法, 绑定一个服务提供者
	IsBind(key string) bool                                        // 关键字凭证是否已经绑定服务提供者
	Make(key string) (interface{}, error)                          // 根据关键字凭证获取一个服务
	MustMake(key string) interface{}                               // 与Make作用一样, 但不返回错误信息
	MakeNew(key string, params []interface{}) (interface{}, error) // 为不同参数启动不同服务
}

type HadeContainer struct {
	Container                            // 强制要求HadeContainer实现Container接口
	providers map[string]ServiceProvider // 存储注册的服务提供者
	instances map[string]interface{}     // 存储具体的实例
	lock      sync.RWMutex               // 读写锁, 这个服务容器是读多于写
}
