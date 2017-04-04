package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/micro/go-micro"
	"fmt"
)

type Config interface {
	Load(filename string)
}

type Configuration struct {
	initialized bool
	Namespace string
	Name      string
	Database  struct {
		Driver         string
		Protocol       string
		Host           string
		Port           uint
		Database       string
		Username       string
		Password       string
		Options        string
		MaxConnections int `yaml:"maxConnections"`
	}
}

func (c *Configuration) Load(filename string) {
	c.initialized = true

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read application.yaml for configuration data.")
	}

	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		panic("Could not unmarshall application.yaml as yaml")
	}
}

func (c Configuration) NewService(version string) micro.Service {
	if !c.initialized {
		panic("Configuration not initialized, call Load() before calling this.")
	}

	return micro.NewService(
		micro.Name(c.Namespace + "." + c.Name),
		micro.Version(version),
	)
}

func (c Configuration) NewConnectionString() string {
	if !c.initialized {
		panic("Configuration not initialized, call Load() before calling this.")
	}

	return c.Database.Username +
		":" +
		c.Database.Password +
		"@" +
		c.Database.Protocol +
		"(" +
		c.Database.Host +
		":" +
		fmt.Sprintf("%d", c.Database.Port) +
		")/" +
		c.Database.Database
}
