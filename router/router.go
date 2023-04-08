package router

import (
	"MyWeb/framework"
	"MyWeb/framework/middleware"
)

func RegisterRoute(c *framework.Core) {
	// 静态路由+HTTP方法匹配
	c.Get("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectAPI := c.Group("/subject")
	{
		// 动态路由
		subjectAPI.Delete("/:id", SubjectDelController)
		subjectAPI.Put("/:id", middleware.Test3(), SubjectUpdateController)
		subjectAPI.Get("/:id", SubjectGetController)
		subjectAPI.Get("/list/all", SubjectListController)

		subjectInnerAPI := subjectAPI.InitGroup("/info")
		{
			subjectInnerAPI.Get("/name", SubjectNameController)
			subjectInnerAPI.Get("/register", SubjectAddController)
		}
	}
}
