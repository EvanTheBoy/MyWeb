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
