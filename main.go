package main

import (
	"log"
	"net/http"

	"github.com/hzy/web/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
