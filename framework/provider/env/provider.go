package env

import (
	"github.com/ascorpio/fly/framework"
	"github.com/ascorpio/fly/framework/contract"
)

type FlyEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *FlyEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewFlyEnv
}

// Boot will called when the service instantiate
func (provider *FlyEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *FlyEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *FlyEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *FlyEnvProvider) Name() string {
	return contract.EnvKey
}
