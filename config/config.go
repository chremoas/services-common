package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"
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

func (c *Configuration) Load(filename string) error {
	c.initialized = true

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("Could not read " + filename + " for configuration data.")
	}

	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return errors.New("Could not unmarshall " + filename + " as yaml")
	}

	return nil
}
