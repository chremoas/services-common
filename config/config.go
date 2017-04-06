package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/micro/go-micro"
)

type Config interface {
	Load(filename string) error
	NewConnectionString() (string, error)
	NewService(version, defaultName string) (micro.Service, error)
	AuthServiceName() (string, error)
}

type Configuration struct {
	initialized bool
	Namespace   string
	Name        string
	Database    struct {
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
	Bot struct {
		BotToken         string `yaml:"botToken"`
		AuthSrvNamespace string `yaml:"authSrvNamespace"`
		DiscordServerId  string `yaml:"discordServerId"`
	}
	OAuth struct {
		ClientId     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
		CallBackUrl  string `yaml:"callBackUrl"`
	} `yaml:"oauth"`
	Net struct {
		ListenHost string `yaml:"listenHost"`
		ListenPort int `yaml:"listenPort"`
	}
	ServiceNames struct {
		AuthSrv string `yaml:"authSrv"`
	} `yaml:"serviceNames"`
}

func (c *Configuration) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("Could not read " + filename + " for configuration data.")
	}

	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return errors.New("Could not unmarshall " + filename + " as yaml")
	}

	// Let's set a default namespace because a lot of people don't care what it actually is
	if c.Namespace == "" {
		c.Namespace = "com.aba-eve"
	}

	c.initialized = true

	return nil
}
