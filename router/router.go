package router

import (
	"MyWeb/framework"
)

func RegisterRoute(c *framework.Core) {
	// 静态路由+HTTP方法匹配
	c.Get("/user/login", UserLoginController)

	// 批量通用前缀
	subjectAPI := c.Group("/subject")
	{
		// 动态路由
		subjectAPI.Delete("/:id", SubjectDelController)
		subjectAPI.Put("/:id", SubjectUpdateController)
		subjectAPI.Get("/:id", SubjectGetController)
		subjectAPI.Get("/list/all", SubjectListController)

		subjectInnerAPI := subjectAPI.InitGroup("/info")
		{
			subjectInnerAPI.Get("/name", SubjectNameController)
			subjectInnerAPI.Get("/register", SubjectAddController)
		}
	}
}
