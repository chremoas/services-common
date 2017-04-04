package config

import (
	"github.com/micro/go-micro"
	"errors"
)

func (c Configuration) NewService(version string) (micro.Service, error) {
	if !c.initialized {
		return nil, errors.New("Configuration not initialized, call Load() before calling this.")
	}

	return micro.NewService(
		micro.Name(c.Namespace + "." + c.Name),
		micro.Version(version),
	), nil
}
