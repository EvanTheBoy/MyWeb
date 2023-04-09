package framework

// IResponse IResponse代表返回方法
type IResponse interface {
	// Json Json输出
	Json(obj interface{}) IResponse

	// Jsonp Jsonp输出
	Jsonp(obj interface{}) IResponse

	// Xml xml输出
	Xml(obj interface{}) IResponse

	// Html html输出
	Html(file string, obj interface{}) IResponse

	// Text string
	Text(format string, values ...interface{}) IResponse

	// Redirect 重定向
	Redirect(path string) IResponse

	// SetHeader header
	SetHeader(key string, val string) IResponse

	// SetCookie Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// SetStatus 设置状态码
	SetStatus(code int) IResponse

	// SetOkStatus 设置200状态
	SetOkStatus() IResponse
}
