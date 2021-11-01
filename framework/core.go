package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {
	getRouter := make(map[string]ControllerHandler)
	postRouter := make(map[string]ControllerHandler)
	putRouter := make(map[string]ControllerHandler)
	deleteRouter := make(map[string]ControllerHandler)

	router := make(map[string]map[string]ControllerHandler)

	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{
		router: router,
	}
}

func (c *Core) GET(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

func (c *Core) POST(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

func (c *Core) PUT(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

func (c *Core) DELETE(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteByRequest(req *http.Request) ControllerHandler {
	uri := req.URL.Path
	log.Println(req.RequestURI)
	method := req.Method
	upperMethod := strings.ToUpper(method)

	if methodHandler, ok := c.router[upperMethod]; ok {
		return methodHandler[uri]
	}
	return nil
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serveHttp")
	ctx := NewContext(r, w)

	router := c.FindRouteByRequest(r)
	if router == nil {
		ctx.Json(http.StatusNotFound, "not found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.Json(http.StatusInternalServerError, err.Error())
	}
}
