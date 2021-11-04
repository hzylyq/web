package framework

import (
	"bytes"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/spf13/cast"
)

type IRequest interface {
	QueryInt(string, int) (int, bool)
	QueryInt64(string, int) (int, bool)
	QueryFloat64(string, float64) (float64, bool)
	QueryFloat32(string, float32) (float32, bool)
	QueryBool(string, bool) (bool, bool)
	QueryString(string, string) (string, bool)
	QueryStringSlice(string, ...string) ([]string, bool)
	Query(string) interface{}

	ParamInt(string, int) (int, bool)
	ParamInt64(string, int) (int, bool)
	ParamFloat64(string, int) (float64, bool)
	ParamFloat32(string, int) (float32, bool)
	ParamBool(string, bool) (bool, bool)
	ParamString(string, string) (string, bool)
	ParamStringSlice(string, ...string) ([]string, bool)
	Param(string) interface{}

	FormInt(string, int) (int, bool)
	FormInt64(string, int) (int, bool)
	FormFloat64(string, int) (float64, bool)
	FormFloat32(string, int) (float64, bool)
	FormBool(string, bool) (bool, bool)
	FormString(string, string) (string, bool)
	FormStringSlice(string, ...string) ([]string, bool)
	FormFile(string) (*multipart.FileHeader, bool)
	Form(string) interface{}

	BindJson(interface{}) error
	BindXml(interface{}) error

	RawData() ([]byte, error)

	Uri() string
	Method() string
	Host() string
	ClientIP() string

	Headers() map[string]string
	Header(string) (string, bool)

	Cookies() map[string]string
	Cookie() string
}

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}``
	return def, false
}

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
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

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def ...string) ([]string, bool) {
	if vals, ok := ctx.QueryAll()[key]; ok {
		return vals, true
	}

	return def, false
}

func (ctx *Context) Query(key string) interface{} {
	if vals := ctx.QueryAll()[key]; len(vals) > 0 {
		return vals[0]
	}
	return nil
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
