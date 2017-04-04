package config

import (
	"testing"
	"github.com/abaeve/auth-common/config"
)

func TestReadConfigFile(t *testing.T) {
	conf := config.Configuration{}
	conf.Load("application.dist.yaml")
	if conf.Namespace != "your.namespace.to.register" {
		t.Error()
	}
	if conf.Name != "yourappname" {
		t.Error()
	}
	if conf.Database.Driver != "mysql" {
		t.Error()
	}
	if conf.Database.Protocol != "tcp" {
		t.Error()
	}
	if conf.Database.Host != "hostname" {
		t.Error()
	}
	if conf.Database.Port != 3306 {
		t.Error()
	}
	if conf.Database.Database != "database" {
		t.Error()
	}
	if conf.Database.Username != "username" {
		t.Error()
	}
	if conf.Database.Password != "password" {
		t.Error()
	}
	if conf.Database.Options != "options" {
		t.Error()
	}
	if conf.Database.MaxConnections != 5 {
		t.Error()
	}
}
