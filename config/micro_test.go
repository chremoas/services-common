package config

import (
	"testing"
	"github.com/micro/go-micro"
)

func TestConfiguration_NewService(t *testing.T) {
	conf := Configuration{}
	conf.Load("application.dist.yaml")
	if _, ok := conf.NewService("1.1.1").(micro.Service); !ok {
		t.Error()
	}
}
