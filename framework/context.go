package framework

import (
	"bytes"
	"context"
	"encoding/binary"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	hasTimeout     bool        // 超时标记
	writerMux      *sync.Mutex // 锁

	handlers []ControllerHandler // 当前请求的handler链条
	index    int                 // 索引, 表示当前请求调用到调用链的哪个节点
	params   map[string]string   // url路由匹配的参数
}

// NewContext 初始化
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// Int64ToBytes 将int64类型的值转换为byte数组
func Int64ToBytes(n int64) []byte {
	byteBuf := bytes.NewBuffer([]byte{})
	err := binary.Write(byteBuf, binary.BigEndian, n)
	if err != nil {
		return nil
	}
	return byteBuf.Bytes()
}

// Next 实现控制链条的核心方法, 获取当前节点的下一个控制器
// 该方法通过维护Context中的下标, 来控制链条移动
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// SetParams 设置路由参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// SetHandlers 设置控制器
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

/*
从这里开始的五个方法, 都是我们要自己封装我们这个结构最基本的方法
*/

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
	return ctx.responseWriter
}

// SetHasTimeout Base, 设置超时时间
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

// HasTimeout Base, 返回有无超时时间的标记
func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

/*
从这里开始，下面的五个方法都是为了实现Context接口, 也就是说, 官方的Context里面
定义了这些方法，需要我们自己去实现
*/

// BaseContext Context, 从这里开始实现标准Context接口
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// Deadline Context, 返回超时的截止时间
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

// Done Context, 被上游树节点通知结束, 返回一个channel
func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

// Err Context, 仅当Done()被完全关闭的时候才返回错误信息解释原因, 否则返回nil
func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

// Value Context, 返回用户的Value
func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}
