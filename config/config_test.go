package config

import (
	"testing"
)

func TestConfiguration_Load(t *testing.T) {
	conf := Configuration{}
	if err := conf.Load("application.dist.yaml"); err != nil {
		t.Errorf("Error from Load(): %s", err)
	}

	if conf.Namespace != "your.namespace.to.register" {
		t.Errorf("Namespace unset")
	}
	if conf.Name != "yourappname" {
		t.Errorf("Name unset")
	}
	if conf.Database.Driver != "mysql" {
		t.Errorf("Database.Driver unset")
	}
	if conf.Database.Protocol != "tcp" {
		t.Errorf("Database.Protocol unset")
	}
	if conf.Database.Host != "hostname" {
		t.Errorf("Database.Host unset")
	}
	if conf.Database.Port != 3306 {
		t.Errorf("Database.Port unset")
	}
	if conf.Database.Database != "database" {
		t.Errorf("Database.Database unset")
	}
	if conf.Database.Username != "username" {
		t.Errorf("Database.Username unset")
	}
	if conf.Database.Password != "password" {
		t.Errorf("Database.Password unset")
	}
	if conf.Database.Options != "options" {
		t.Errorf("Database.Options unset")
	}
	if conf.Database.MaxConnections != 5 {
		t.Errorf("Database.MaxConnections unset")
	}
}

func TestConfiguration_Load_NoFile(t *testing.T) {
	conf := Configuration{}
	if err := conf.Load("application.nofile.yaml"); err == nil {
		t.Errorf("No error from Load() with no file")
	}
}

func TestConfiguration_Load_InvalidFile(t *testing.T) {
	conf := Configuration{}
	if err := conf.Load("application.invalid.yaml"); err == nil {
		t.Errorf("No error from Load() with invalid file")
	}
}
