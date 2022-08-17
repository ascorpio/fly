package kernel

import (
	"github.com/ascorpio/fly/framework/gin"
	"net/http"
)

// 引擎服务
type FlyKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewFlyKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &FlyKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *FlyKernelService) HttpEngine() http.Handler {
	return s.engine
}
