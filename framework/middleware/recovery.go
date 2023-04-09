package middleware

import "MyWeb/framework"

// Recovery 将协程中的函数异常进行捕获
func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// 捕获c.Next()中出现的panic异常
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()
		// 调用Next执行具体的业务
		err1 := c.Next()
		if err1 != nil {
			return err1
		}
		return nil
	}
}
