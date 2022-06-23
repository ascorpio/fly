package main

import (
	"go_frame/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		// 自定义请求核心处理函数
		Handler: framework.NewCore(),
		// 请求监听地址
		Addr: ":8888",
	}
	server.ListenAndServe()
}
