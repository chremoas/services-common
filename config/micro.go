package config

import (
	"errors"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var configurationFile string

func NewService(version string, initFunc func(configuration *Configuration) error) micro.Service {
	service := micro.NewService(
		micro.Version(version),
		micro.BeforeStart(
			func() error {
				configuration := Configuration{}

				configuration.Load(configurationFile)

				if !configuration.initialized {
					return errors.New("Configuration not initialized, call Load() before calling this.")
				}

				return initFunc(&configuration)
			},
		),
		micro.Flags(
			cli.StringFlag{
				Name:        "configuration_file",
				Usage:       "The yaml configuration file for the service being loaded",
				Value:       "/etc/auth-srv/application.yaml",
				EnvVar:      "CONFIGURATION_FILE",
				Destination: &configurationFile,
			},
		),
	)

	service.Init()

	return service
}

func NilInit(configuration *Configuration) error {
	return nil
}
