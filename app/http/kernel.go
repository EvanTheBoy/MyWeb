package http

import "github.com/gohade/my-web/framework/gin"

// NewHttpEngine is command
func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 业务绑定路由操作
	Routes(r)
	return r, nil
}
