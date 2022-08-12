package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Context 自定义
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if val, ok := params[key]; ok {
		vLen := len(val)
		if vLen > 0 {
			intVal, err := strconv.Atoi(val[vLen-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if val, ok := params[key]; ok {
		vLen := len(val)
		if vLen > 0 {
			return val[vLen-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if val, ok := params[key]; ok {
		return val
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) FromInt(key string, def int) int {
	params := ctx.FormAll()
	if val, ok := params[key]; ok {
		vLen := len(val)
		if vLen > 0 {
			intVal, err := strconv.Atoi(val[vLen-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (ctx *Context) FromString(key string, def string) string {
	params := ctx.FormAll()
	if val, ok := params[key]; ok {
		vLen := len(val)
		if vLen > 0 {
			return val[vLen-1]
		}
	}
	return def
}

func (ctx *Context) FromArray(key string, def []string) []string {
	params := ctx.FormAll()
	if val, ok := params[key]; ok {
		return val
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) BindJson(obj any) error {
	if ctx.request != nil {
		body, err := io.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = io.NopCloser(bytes.NewBuffer(body))
		return json.Unmarshal(body, obj)
	} else {
		return errors.New("ctx.request empty")
	}
}

func (ctx *Context) Json(status int, obj any) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj any, template template.Template) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
