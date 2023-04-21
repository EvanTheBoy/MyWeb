package contract

import "net/http"

// KernelKey 提供 kernel 服务凭证
const KernelKey = "hade:kernel"

// Kernel 提供kernel为核心的结构
type Kernel interface {
	// HttpEngine 提供gin的Engine结构
	HttpEngine() http.Handler
}
