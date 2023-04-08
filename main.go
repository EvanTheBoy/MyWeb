package main

import (
	"MyWeb/framework"
	"MyWeb/framework/middleware"
	"MyWeb/router"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()
	core.RegisterMiddleware(middleware.Recovery())
	core.RegisterMiddleware(middleware.Cost())
	core.RegisterMiddleware(middleware.Timeout(1 * time.Second))
	router.RegisterRoute(core)
	server := http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
