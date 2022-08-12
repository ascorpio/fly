package framework

import "net/http"

// Core 框架核心
type Core struct {
	router map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// 框架核心结构实现 Handle 接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
