package router

import (
	"MyWeb/app"
	"MyWeb/framework"
)

func RegisterRoute(c *framework.Core) {
	// 静态路由+HTTP方法匹配
	c.Get("/user/login", app.UserLoginController)

	// 批量通用前缀
	subjectAPI := c.Group("/subject")
	{
		// 动态路由
		subjectAPI.Delete("/:id", app.SubjectDelController)
		subjectAPI.Put("/:id", app.SubjectUpdateController)
		subjectAPI.Get("/:id", app.SubjectGetController)
		subjectAPI.Get("/list/all", app.SubjectListController)

		subjectInnerAPI := subjectAPI.InitGroup("/info")
		{
			subjectInnerAPI.Get("/name", app.SubjectNameController)
		}
	}
}
