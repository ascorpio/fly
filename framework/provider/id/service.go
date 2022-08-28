package id

import (
	"github.com/rs/xid"
)

type FlyIDService struct {
}

func NewFlyIDService(params ...interface{}) (interface{}, error) {
	return &FlyIDService{}, nil
}

func (s *FlyIDService) NewID() string {
	return xid.New().String()
}
