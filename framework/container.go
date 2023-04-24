package framework

import (
	"errors"
	"fmt"
	"sync"
)

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

// NewHadeContainer 初始化一个服务容器
func NewHadeContainer() *HadeContainer {
	return &HadeContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出注册的服务容器中的关键字
func (hade *HadeContainer) PrintProviders() []string {
	var ret []string
	for _, provider := range hade.providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做绑定
func (hade *HadeContainer) Bind(provider ServiceProvider) error {
	// 写操作, 先加锁
	hade.lock.Lock()
	defer hade.lock.Unlock()

	key := provider.Name()
	hade.providers[key] = provider

	// 判断是否在注册的时候实例化服务, 为false那么现在就要实例化
	if provider.IsDefer() == false {
		// 在实例化前做一些准备工作
		if err := provider.Boot(hade); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(hade)
		method := provider.Register(hade)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		hade.instances[key] = instance
	}
	return nil
}

// IsBind 查询是否绑定了服务提供者
func (hade *HadeContainer) IsBind(key string) bool {
	return hade.findServiceProvider(key) != nil
}

// 查找是否存在这个服务的提供者
func (hade *HadeContainer) findServiceProvider(key string) ServiceProvider {
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	if sp, ok := hade.providers[key]; ok {
		return sp
	}
	return nil
}

// 创建一个新的服务实例
func (hade *HadeContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// 先做一些准备工作
	if err := sp.Boot(hade); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(hade)
	}
	// 创建新建一个实例的方法
	method := sp.Register(hade)
	instance, err := method(hade)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return instance, err
}

// 真正的实例化一个服务
func (hade *HadeContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	// 上锁
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	// 先检查是否已经注册了这个服务
	sp := hade.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not registered")
	}
	// 这个不为空, 那就是MakeNew方法
	if forceNew {
		return hade.newInstance(sp, params)
	}
	// 若容器中已经实例化了, 那就拿来直接用
	if ins, ok := hade.instances[key]; ok {
		return ins, nil
	}
	// 容器中还未实例化, 那就实例化吧
	inst, err := hade.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	hade.instances[key] = inst
	return inst, nil
}

func (hade *HadeContainer) Make(key string) (interface{}, error) {
	return hade.make(key, nil, false)
}

// MakeNew 在获取服务实例的时候, 按照不同参数进行初始化
func (hade *HadeContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return hade.make(key, params, true)
}

func (hade *HadeContainer) MustMake(key string) interface{} {
	serve, err := hade.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serve
}
