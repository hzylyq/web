package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := make(map[string]*Tree)

	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{
		router: router,
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) GET(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router err:", err)
	}
}

func (c *Core) POST(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router err:", err)
	}
}

func (c *Core) PUT(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router err:", err)
	}
}

func (c *Core) DELETE(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router err:", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteByRequest(req *http.Request) []ControllerHandler {
	uri := req.URL.Path
	method := req.Method
	upperMethod := strings.ToUpper(method)

	if methodHandler, ok := c.router[upperMethod]; ok {
		return methodHandler.FindHandler(uri)
	}
	return nil
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serveHttp")
	ctx := NewContext(r, w)

	handlers := c.FindRouteByRequest(r)
	if len(handlers) == 0 {
		ctx.Json(http.StatusNotFound, "not found")
		return
	}

	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(http.StatusInternalServerError, "internal err")
		return
	}
}
