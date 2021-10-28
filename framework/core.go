package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: make(map[string]ControllerHandler),
	}
}

func (c *Core) Set(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serveHttp")
	ctx := NewContext(r, w)

	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)
}
