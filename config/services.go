package config

import "errors"

// Helper to resolve the auth-srv service to whatever the consumer has set things to
// I'm fan of being able to personalize things to a certain extent if they really want to.
// Who knows, we may want to attempt a multi tenant situation where multiple different services
// live in a weird configuration of same or different namespaces and same or different names
// within a shared namespace.  This give us that flexibility while also allowing us to severly shoot
// ourselves in the foot...
func (c Configuration) AuthServiceName() (string, error) {
	if !c.initialized {
		return "", errors.New("Configuration not initialized, call Load() before calling this.")
	}

	return c.Namespace + "." + c.ServiceNames.AuthSrv, nil
}
