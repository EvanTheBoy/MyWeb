package router

import (
	"github.com/gohade/my-web/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	// 等待10s才结束
	time.Sleep(10 * time.Second)
	// 输出结果
	c.ISetOkStatus().IJson("ok, UserLoginController:" + foo)
}
