package framework

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
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

/*
接下来的几个函数都是为了实现Context里面基本的request请求,
即, 用Query来获取对应的get参数
*/

// QueryAll 在内部我们调用的是Query函数, 这个函数会解析"raw查询", 然后返回解析后的值
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		// Query函数返回的类型就是这个QueryAll函数返回的类型
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// QueryInt 查询并返回整数型的值
func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			intVal, err := strconv.Atoi(values[length-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

// QueryString 查询并返回字符串类型的值
func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			return values[length-1]
		}
	}
	return def
}

// QueryArray 查询并返回一个字符串类型的数组的值
func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

/*
这一部分是向服务端发送数据, 使用的是post。
*/

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			intVal, err := strconv.Atoi(values[length-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		length := len(values)
		if length > 0 {
			return values[length-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		n, err := io.Copy(os.Stdout, ctx.request.Body)
		if err != nil {
			return err
		}
		// 使用io.Copy()更加高效, 但因为这个函数返回的是一个int64类型的值, 需要转换为byte数组
		body := Int64ToBytes(n)
		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request is empty")
	}
	return nil
}

/*
以下为response部分的函数代码
*/

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.WriteHeader(status)
	_, err1 := ctx.responseWriter.Write(byt)
	if err1 != nil {
		return err
	}
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
