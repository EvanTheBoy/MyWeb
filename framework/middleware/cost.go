package middleware

import (
	"MyWeb/framework"
	"log"
	"time"
)

// Cost 记录耗时
func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// 记录开始时间
		start := time.Now()
		// 调用具体的业务逻辑
		err := c.Next()
		if err != nil {
			return err
		}
		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api url:%v, cost:%v", c.GetRequest().RequestURI, cost.Seconds())
		return nil
	}
}
