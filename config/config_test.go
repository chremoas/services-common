package config

import (
	"fmt"
	"testing"
	"github.com/abaeve/auth-common/config"
)

func TestReadConfigFile(t *testing.T) {
	conf := Configuration{}
	conf.load('application.dist.yaml')
	fmt.Print(conf)
}
