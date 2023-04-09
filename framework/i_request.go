package framework

import "mime/multipart"

// IRequest 里面含有请求包含的方法
type IRequest interface {
	// QueryInt 及以下函数, 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// ParamInt 及以下函数, 路由匹配中带的参数
	// 形如 /book/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	// FormInt 及以下函数, 都是form表单中带的参数
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

	// BindJson json body
	BindJson(obj interface{}) error

	// BindXml xml body
	BindXml(obj interface{}) error

	// GetRawData 其他格式
	GetRawData() ([]byte, error)

	// Uri 及以下函数, 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIP() string

	// Headers Header header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// Cookies Cookie cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}
