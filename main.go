package main

import (
	"log"
	"net/http"

	"github.com/hzy/web/framework"
	"github.com/hzy/web/framework/middleware"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)

	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
