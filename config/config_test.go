package config

import (
	"fmt"
	"testing"
	"github.com/abaeve/auth-common/config"
)

func TestReadConfigFile(t *testing.T) {
	conf := config.Configuration{}
	conf.Load("application.dist.yaml")
	fmt.Print(conf)
}
