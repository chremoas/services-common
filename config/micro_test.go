package config

import (
	"github.com/micro/go-micro"
	"testing"
)

func TestConfiguration_NewService(t *testing.T) {
	conf := Configuration{}
	conf.Load("application.dist.yaml")
	service, err := conf.NewService("1.1.1", "")
	if _, ok := service.(micro.Service); !ok {
		t.Errorf("service is not a micro.Service!")
	}
	if err != nil {
		t.Errorf("Error not set to nil, message is: %s", err)
	}
}

func TestConfiguration_NewService_NoName(t *testing.T) {
	conf := Configuration{}
	conf.Load("application.noname.yaml")
	service, err := conf.NewService("1.1.1", "defaultname")
	if _, ok := service.(micro.Service); !ok {
		t.Errorf("service is not a micro.Service!")
	}
	if err != nil {
		t.Errorf("Error not set to nil, message is: %s", err)
	}
}

func TestConfiguration_NewService_NoName_NoDefault(t *testing.T) {
	conf := Configuration{}
	conf.Load("application.noname.yaml")
	service, err := conf.NewService("1.1.1", "")
	if _, ok := service.(micro.Service); ok {
		t.Errorf("service is not a micro.Service!")
	}
	if err == nil {
		t.Errorf("Error not set to nil, message is: %s", err)
	}
}

func TestConfiguration_NewService_NoConfLoaded(t *testing.T) {
	conf := Configuration{}
	service, err := conf.NewService("1.1.1", "")
	if err == nil {
		if _, ok := service.(micro.Service); ok {
			t.Error("Error is nil and yet we have a valid service?")
		}
		t.Error("Error set to nil, at least we don't have a valid service")
	}
}
