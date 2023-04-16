package demo

import (
	"fmt"
	"github.com/gohade/my-web/framework"
)

type DemoService struct {
	Service
	c framework.Container
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "I am foo",
	}
}

// NewDemoService 初始化服务实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	fmt.Println("New demo service")
	return &DemoService{c: c}, nil
}
