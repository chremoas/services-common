package config

import "github.com/micro/go-micro"

func (c Configuration) NewService(version string) micro.Service {
	if !c.initialized {
		panic("Configuration not initialized, call Load() before calling this.")
	}

	return micro.NewService(
		micro.Name(c.Namespace + "." + c.Name),
		micro.Version(version),
	)
}
