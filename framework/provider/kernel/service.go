package kernel

import (
	"github.com/gohade/my-web/framework/gin"
	"net/http"
)

// HadeKernelService 引擎服务
type HadeKernelService struct {
	engine *gin.Engine
}

func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &HadeKernelService{
		engine: httpEngine,
	}, nil
}

func (s *HadeKernelService) HttpEngine() http.Handler {
	return s.engine
}
