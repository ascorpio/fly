package framework

import (
	"context"
	"html/template"
	"net/http"
)

// Context 自定义
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
}

func NewContext(r *http.Request, w http.ResponseWriter) {}
func (ctx *Context) WriterMux()                         {}
func (ctx *Context) GetRequest()                        {}
func (ctx *Context) GetResponse()                       {}
func (ctx *Context) SetHasTimeout()                     {}
func (ctx *Context) HasTimeout()                        {}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() {}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err()          {}
func (ctx *Context) Value(key any) {}

func (ctx *Context) QueryInt(key string, def int)        {}
func (ctx *Context) QueryString(key string, def string)  {}
func (ctx *Context) QueryArray(key string, def []string) {}
func (ctx *Context) QueryAll()                           {}
func (ctx *Context) FromInt(key string, def int)         {}
func (ctx *Context) FromString(key string, def string)   {}
func (ctx *Context) FromArray(key string, def []string)  {}
func (ctx *Context) FromAll()                            {}
func (ctx *Context) BindJson(obj any)                    {}

func (ctx *Context) Json(status int, obj any) error {
	// 待实现
	return nil
}
func (ctx *Context) HTML(status int, obj any, template template.Template) {}
func (ctx *Context) Text(status int, obj string)                          {}
