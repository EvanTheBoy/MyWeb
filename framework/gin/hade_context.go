// Package gin Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
	"github.com/gohade/my-web/framework"
)

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}

/*
以下, 重写这些函数, 是因为我们要把服务容器融合进框架中
*/

func (c *Context) Make(key string) (interface{}, error) {
	return c.container.Make(key)
}

func (c *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return c.container.MakeNew(key, params)
}

func (c *Context) MustMake(key string) interface{} {
	return c.container.MustMake(key)
}

func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}
