package main

import (
	"fly/framework"
	"net/http"
)

func main() {
	// 设置路由
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// 自定义请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}
