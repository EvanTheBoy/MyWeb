package main

import (
	"MyWeb/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	server := http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
