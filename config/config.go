package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config interface {
	Load(filename string)
}

type Configuration struct {
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
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic("Could not read application.yaml for configuration data.")
	}

	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		panic("Could not unmarshall application.yaml as yaml")
	}
}
