package main

import (
	"MyWeb/framework"
	"MyWeb/router"
	"net/http"
)

func main() {
	core := framework.NewCore()
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
