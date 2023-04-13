package router

import (
	"github.com/gohade/my-web/framework/gin"
	"github.com/gohade/my-web/framework/middleware"
)

func RegisterRoute(c *gin.Engine) {
	// 静态路由+HTTP方法匹配
	c.GET("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectAPI := c.Group("/subject")
	{
		// 动态路由
		subjectAPI.DELETE("/:id", SubjectDelController)
		subjectAPI.PUT("/:id", middleware.Test1(), SubjectUpdateController)
		subjectAPI.GET("/:id", SubjectGetController)
		subjectAPI.GET("/list/all", SubjectListController)

		subjectInnerAPI := subjectAPI.Group("/info")
		{
			subjectInnerAPI.GET("/name", SubjectNameController)
			subjectInnerAPI.GET("/register", SubjectAddController)
		}
	}
}
