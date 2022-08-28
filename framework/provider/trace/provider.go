package trace

import (
	"github.com/ascorpio/fly/framework"
	"github.com/ascorpio/fly/framework/contract"
)

type FlyTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *FlyTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewFlyTraceService
}

// Boot will called when the service instantiate
func (provider *FlyTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *FlyTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *FlyTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *FlyTraceProvider) Name() string {
	return contract.TraceKey
}
