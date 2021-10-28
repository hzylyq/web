package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	req  *http.Request
	resp http.ResponseWriter
	ctx  context.Context

	hasTimeOut bool
	writeMux   *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		req:      r,
		resp:     w,
		ctx:      r.Context(),
		writeMux: &sync.Mutex{},
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

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			intVal, err := strconv.Atoi(vals[len(vals)-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) QueryString(key, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[len(vals)-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string) []string {
	return ctx.QueryAll()[key]
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.req != nil {
		return ctx.req.URL.Query()
	}

	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			intVal, err := strconv.Atoi(vals[len(vals)-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) FormString(key, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[len(vals)-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string) []string {
	return ctx.FormAll()[key]
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx != nil {
		return ctx.req.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.req == nil {
		return errors.New("ctx.request empty")
	}

	body, err := io.ReadAll(ctx.req.Body)
	if err != nil {
		return err
	}

	ctx.req.Body = io.NopCloser(bytes.NewBuffer(body))

	return json.Unmarshal(body, obj)
}

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}

	ctx.resp.Header().Set("Content-Type", "application/json")
	ctx.resp.WriteHeader(status)
	b, err := json.Marshal(obj)
	if err != nil {
		ctx.resp.WriteHeader(500)
		return err
	}
	_, err = ctx.resp.Write(b)
	return err
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
