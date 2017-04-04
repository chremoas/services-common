package config

import (
	"testing"
	"github.com/micro/go-micro"
)

func TestConfiguration_NewService(t *testing.T) {
	conf := Configuration{}
	conf.Load("application.dist.yaml")
	service, err := conf.NewService("1.1.1")
	if _, ok := service.(micro.Service); !ok {
		t.Errorf("service is not a micro.Service!")
	}
	if err != nil {
		t.Errorf("Error not set to nil, message is: %s", err)
	}
}

func TestConfiguration_NewService_NoConfLoaded(t *testing.T) {
	conf := Configuration{}
	service, err := conf.NewService("1.1.1")
	if err == nil {
		if _, ok := service.(micro.Service); ok {
			t.Errorf("Error is nil and yet we have a valid service?")
		}
		t.Errorf("Error set to nil, at least we don't have a valid service")
	}
}
