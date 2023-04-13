package middleware

import (
	"github.com/gohade/my-web/framework/gin"
	"log"
	"time"
)

// Cost 记录耗时
func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()
		log.Printf("api uri start: %v", c.Request.RequestURI)
		// 执行具体的业务逻辑
		c.Next()

		// 记录结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api url:%v, cost:%v", c.Request.RequestURI, cost.Seconds())
	}
}
