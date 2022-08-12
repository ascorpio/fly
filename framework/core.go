package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心
type Core struct {
	router map[string]*Tree
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

func (c *Core) Get(url string, handler ControllerHandler) {
	c.addRouter("GET", url, handler)
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.addRouter("POST", url, handler)
}

func (c *Core) Put(url string, handler ControllerHandler) {
	c.addRouter("PUT", url, handler)
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	c.addRouter("DELETE", url, handler)
}

func (c *Core) addRouter(method, url string, handler ControllerHandler) {
	if err := c.router[method].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// 框架核心结构实现 Handle 接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 自定义 Context
	ctx := NewContext(request, response)

	// 一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.FindRouteByRequest(request)
	if router == nil {
		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

// FindRouteByRequest 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
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
