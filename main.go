package main

import (
	"encoding/json"
	"go_frame/framework"
	"net/http"
	"strconv"
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

// 控制器
func Foo1(request *http.Request, response http.ResponseWriter) {
	obj := map[string]any{
		"data": nil,
	}
	// 设置控制器 response 的 header 部分
	response.Header().Set("Content-Type", "application/json")

	// 从请求体中获取参数
	foo := request.PostFormValue("foo")
	if foo == "" {
		foo = "10"
	}
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回结构
	obj["data"] = fooInt
	byt, err := json.Marshal(obj)
	if err != nil {
		response.WriteHeader(500)
		return
	}
	// 构建返回状态，输出返回结构
	response.WriteHeader(200)
	response.Write(byt)
	return
}
