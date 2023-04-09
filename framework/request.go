package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io"
	"mime/multipart"
	"os"
)

const (
	defaultMultipartMemory = 32 << 20
)

// QueryAll 获取请求地址中所有的参数
// 在内部我们调用的是Query函数, 这个函数会解析"raw查询", 然后返回解析后的值
func (ctx *Context) QueryAll() map[string][]string {
	if ctx != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// QueryInt 获取int类型的请求参数
func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

// QueryString 获取字符串类型的请求参数
func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToString(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values, ok
	}
	return def, false
}

func (ctx *Context) Query(key string) interface{} {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values[0]
	}
	return nil
}

// Param 获取路由参数
func (ctx *Context) Param(key string) interface{} {
	if ctx.params != nil {
		if values, ok := ctx.params[key]; ok {
			return values[0]
		}
	}
	return nil
}

// ParamInt 路由匹配中携带的参数
// 形如 /book/:id
func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if val := ctx.Param(key); val != nil {
		// 使用cast进行转换
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.Param(key); val != nil {
		// 使用cast进行转换
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.Param(key); val != nil {
		// 使用cast进行转换
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.Param(key); val != nil {
		// 使用cast进行转换
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := ctx.Param(key); val != nil {
		// 使用cast进行转换
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if val := ctx.Param(key); val != nil {
		// 使用cast进行转换
		return cast.ToString(val), true
	}
	return def, false
}

// FormAll 这部分是解析表单数据
func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		err := ctx.request.ParseForm()
		if err != nil {
			return nil
		}
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return cast.ToString(values[0]), true
	}
	return def, false
}

func (ctx *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values, true
	}
	return def, false
}

func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {
	if ctx.request.MultipartForm == nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.request.FormFile(key)
	if err != nil {
		return nil, err
	}
	err1 := f.Close()
	if err1 != nil {
		return nil, err1
	}
	return fh, err
}

func (ctx *Context) Form(key string) interface{} {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}
	return nil
}

// BindJson 将body文本解析到obj结构体中
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		// 读取文本
		n, err := io.Copy(os.Stdout, ctx.request.Body)
		if err != nil {
			return err
		}
		// 使用io.Copy()更加高效, 但因为这个函数返回的是一个int64类型的值, 需要转换为byte数组
		body := Int64ToBytes(n)
		// 重新填充request.Body, 为后续的逻辑二次读取做准备
		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))
		// 解析到obj结构体中
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request is empty")
	}
	return nil
}

// BindXml xml body
func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.request != nil {
		n, err := io.Copy(os.Stdout, ctx.request.Body)
		if err != nil {
			return err
		}
		body := Int64ToBytes(n)
		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))
		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// GetRawData 其他格式
func (ctx *Context) GetRawData() ([]byte, error) {
	if ctx.request != nil {
		n, err := io.Copy(os.Stdout, ctx.request.Body)
		if err != nil {
			return nil, err
		}
		body := Int64ToBytes(n)
		// 填充回去, 方便下次再来取
		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil, errors.New("ctx.request empty")
}

// Uri Method Host 都属于基础信息
func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) Host() string {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIP() string {
	r := ctx.request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// Headers header
func (ctx *Context) Headers() map[string][]string {
	return ctx.request.Header
}

func (ctx *Context) Header(key string) (string, bool) {
	values := ctx.request.Header.Values(key)
	if values == nil || len(values) <= 0 {
		return "", false
	}
	return values[0], true
}

// Cookies cookie
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

func (ctx *Context) Cookie(key string) (string, bool) {
	cookies := ctx.Cookies()
	if val, ok := cookies[key]; ok {
		return val, true
	}
	return "", false
}
