package config

import (
	"errors"
	"github.com/micro/go-micro"
)

func (c Configuration) NewService(version, defaultName string) (micro.Service, error) {
	if !c.initialized {
		return nil, errors.New("Configuration not initialized, call Load() before calling this.")
	}

	if c.Name == "" {
		if defaultName == "" {
			return nil, errors.New("Name not set in yaml config or as a default value")
		}
		c.Name = defaultName
	}

	return micro.NewService(
		micro.Name(c.Namespace+"."+c.Name),
		micro.Version(version),
	), nil
}
