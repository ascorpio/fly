package id

import (
	"github.com/ascorpio/fly/framework"
	"github.com/ascorpio/fly/framework/contract"
)

type FlyIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *FlyIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewFlyIDService
}

// Boot will called when the service instantiate
func (provider *FlyIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *FlyIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *FlyIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *FlyIDProvider) Name() string {
	return contract.IDKey
}
