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

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:   r,
		response:  w,
		ctx:       r.Context(),
		writerMux: &sync.Mutex{},
	}
}
