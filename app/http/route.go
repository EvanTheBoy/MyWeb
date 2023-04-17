package http

import (
	"github.com/gohade/my-web/app/http/module/demo"
	"github.com/gohade/my-web/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
