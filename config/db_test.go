package config

import (
	"testing"
)

func TestConfiguration_NewConnectionString(t *testing.T) {
	conf := Configuration{}
	conf.Load("application.dist.yaml")
	DSN := conf.NewConnectionString()
	if DSN != "username:password@tcp(hostname:3306)/database" {
		t.Error()
	}
}
