package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Application struct {
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
}

func (c *Configuration) load(filename string) bool {
	data, err := ioutil.ReadFile(filename)

	//TODO: Candidate for shared function for all my services.
	if err != nil {
		panic("Could not read application.yaml for configuration data.")
	}

	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		panic("Could not unmarshall application.yaml as yaml")
	}

	return true
}
