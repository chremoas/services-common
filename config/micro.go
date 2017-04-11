package config

import (
	"errors"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var configurationFile string

// Builds and inits a new micro.Service object for use.  The initFunc functor being asked for will be inserted
// into the services options as a BeforeStart which will be called DURING the service.Run invocation but BEFORE the
// service is fully up and operational.  All of your initialization code that you need should go into this initFunc.
// If you don't need init code then feel free to use the NilInit function exported out of this package.
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
