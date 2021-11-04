package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	req  *http.Request
	resp http.ResponseWriter
	ctx  context.Context

	hasTimeOut bool
	writeMux   *sync.Mutex

	handlers []ControllerHandler
	index    int
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		req:      r,
		resp:     w,
		ctx:      r.Context(),
		writeMux: &sync.Mutex{},
		index:    -1,
	}
}

func (ctx *Context) WriteMux() *sync.Mutex {
	return ctx.writeMux
}

func (ctx *Context) Request() *http.Request {
	return ctx.req
}

func (ctx *Context) Response() http.ResponseWriter {
	return ctx.resp
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeOut = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeOut
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		return ctx.handlers[ctx.index](ctx)
	}

	return nil
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.req.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Deadline() (time.Time, bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}
