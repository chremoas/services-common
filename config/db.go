package config

import "fmt"

func (c Configuration) NewConnectionString() string {
	if !c.initialized {
		panic("Configuration not initialized, call Load() before calling this.")
	}

	return c.Database.Username +
		":" +
		c.Database.Password +
		"@" +
		c.Database.Protocol +
		"(" +
		c.Database.Host +
		":" +
		fmt.Sprintf("%d", c.Database.Port) +
		")/" +
		c.Database.Database
}
