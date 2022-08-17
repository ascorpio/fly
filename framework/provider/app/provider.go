package app

import (
	"github.com/ascorpio/fly/framework"
	"github.com/ascorpio/fly/framework/contract"
)

// FlyAppProvider 提供App的具体实现方法
type FlyAppProvider struct {
	BaseFolder string
}

// Register 注册FlyApp方法
func (h *FlyAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewFlyApp
}

// Boot 启动调用
func (h *FlyAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *FlyAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *FlyAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *FlyAppProvider) Name() string {
	return contract.AppKey
}
