package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心
type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{
		"GET":    NewTree(),
		"POST":   NewTree(),
		"PUT":    NewTree(),
		"DELETE": NewTree(),
	}
	return &Core{
		router: router,
	}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	c.addRouter("GET", url, handlers...)
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	c.addRouter("POST", url, handlers...)
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	c.addRouter("PUT", url, handlers...)
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	c.addRouter("DELETE", url, handlers...)
}

func (c *Core) addRouter(method, url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router[method].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 框架核心结构实现 Handle 接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 自定义 Context
	ctx := NewContext(request, response)

	// 寻找路由
	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		ctx.Json(404, "not found")
		return
	}

	// 设置 context 中的 handlers 字段
	ctx.SetHandlers(handlers)

	// 调用路由函数，如果返回 err 代表内部错误，返回 500 状态
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	uri := request.URL.Path
	upperMethod := strings.ToUpper(request.Method)
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// Group 初始化分组
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// Use 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}
