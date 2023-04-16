package app

import (
	"errors"
	"github.com/gohade/my-web/framework"
)

// HadeApp 代表Hade框架的app实现
type HadeApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

// NewHadeApp 初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &HadeApp{
		container:  container,
		baseFolder: baseFolder,
	}, nil
}

// Version 实现版本
func (h HadeApp) Version() string {
	return "0.0.1"
}
