package framework

import (
	"context"
	"net/http"
	"sync"
)

type Context struct {
	request    *http.Request
	response   http.ResponseWriter
	ctx        context.Context
	handler    ControllerHandler
	hasTimeout bool
	writerMux  *sync.Mutex
}

// NewContext 初始化
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:   r,
		response:  w,
		ctx:       r.Context(),
		writerMux: &sync.Mutex{},
	}
}

// WriteMux Base, 属于基本的函数功能
func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writerMux
}

// GetRequest Base, 获取请求
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

// GetResponse Base, 获取响应
func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.response
}

// SetHasTimeout Base, 设置超时时间
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

// HasTimeout Base, 返回有无超时时间的标记
func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// BaseContext Context, 从这里开始实现标准Context接口
func (ctx *Context) BaseContext() *Context {
	return ctx.BaseContext()
}

// Deadline Context, 返回超时的截止时间
func (ctx *Context) Deadline() *Context {
	return ctx.Deadline()
}

// Done Context, 被上游树节点通知结束
func (ctx *Context) Done() *Context {
	return ctx.Done()
}
