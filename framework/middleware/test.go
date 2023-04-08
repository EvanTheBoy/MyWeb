package middleware

import (
	"MyWeb/framework"
	"fmt"
)

func Test1() framework.ControllerHandler {
	// 使用回调函数
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		err1 := c.Next()
		if err1 != nil {
			return err1
		}
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	// 使用回调函数
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		err1 := c.Next()
		if err1 != nil {
			return err1
		}
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	// 使用回调函数
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		err1 := c.Next()
		if err1 != nil {
			return err1
		}
		fmt.Println("middleware post test3")
		return nil
	}
}
