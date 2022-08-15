package main

import (
	"context"
	"github.com/ascorpio/fly/framework/gin"
	"github.com/ascorpio/fly/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 启动 Goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号： SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 阻塞当前 Goroutine 等待信号
	<-quit

	// 最多5秒，超过后强制进行关闭
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 调用 server.Shutdown graceful 结束
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
