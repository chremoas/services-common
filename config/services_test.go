package config

import "testing"

func TestConfiguration_AuthServiceName(t *testing.T) {
	config := Configuration{}

	config.Load("application.dist.yaml")

	expectedAuthSrvName := "your.namespace.to.register.auth-srv"

	if authSrv, err := config.AuthServiceName(); authSrv != expectedAuthSrvName || err != nil {
		t.Errorf("authSrv: (%s) but should have been: (%s), err: (%+v)", authSrv, expectedAuthSrvName, err)
	}
}
