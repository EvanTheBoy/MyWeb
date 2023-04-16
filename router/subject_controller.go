package router

import (
	"github.com/gohade/my-web/app/provider/demo"
	"github.com/gohade/my-web/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

func SubjectListController(c *gin.Context) {
	// 获取demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)
	// 调用服务实例的方法
	foo := demoService.GetFoo()
	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectGetController")
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
