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
		t.Error("Namespace unset")
	}
	if conf.Name != "yourappname" {
		t.Error("Name unset")
	}
	if conf.Database.Driver != "mysql" {
		t.Error("Database.Driver unset")
	}
	if conf.Database.Protocol != "tcp" {
		t.Error("Database.Protocol unset")
	}
	if conf.Database.Host != "hostname" {
		t.Error("Database.Host unset")
	}
	if conf.Database.Port != 3306 {
		t.Error("Database.Port unset")
	}
	if conf.Database.Database != "database" {
		t.Error("Database.Database unset")
	}
	if conf.Database.Username != "username" {
		t.Error("Database.Username unset")
	}
	if conf.Database.Password != "password" {
		t.Error("Database.Password unset")
	}
	if conf.Database.Options != "options" {
		t.Error("Database.Options unset")
	}
	if conf.Database.MaxConnections != 5 {
		t.Error("Database.MaxConnections unset")
	}
	if conf.OAuth.ClientId != "client_id" {
		t.Error("OAuth.ClientId unset")
	}
	if conf.OAuth.ClientSecret != "client_secret" {
		t.Error("OAuth.ClientSecret unset")
	}
	if conf.OAuth.CallBackUrl != "callback_url" {
		t.Error("OAuth.CallBackUrl unset")
	}
	if conf.Net.ListenHost != "localhost" {
		t.Error("Net.ListenHost unset")
	}
	if conf.Net.ListenPort != 80 {
		t.Error("Net.ListenPort unset")
	}
	if conf.ServiceNames.AuthSrv != "auth-srv" {
		t.Error("ServiceNames.AuthSrv unset")
	}
	if conf.Bot.DiscordServerId != "You bot token here, do not prepend Bot... we'll do that for you" {
		t.Error("Bot.DiscordServerId unset")
	}
	if conf.Bot.AuthSrvNamespace != "namespace that the auth-srv instance lives in" {
		t.Error("Bot.AuthSrvNamespace unset")
	}
	if conf.Bot.BotToken != "You bot token here, do not prepend Bot... we'll do that for you" {
		t.Error("Bot.BotToken unset")
	}
}

func TestConfiguration_Load_NoFile(t *testing.T) {
	conf := Configuration{}
	if err := conf.Load("application.nofile.yaml"); err == nil {
		t.Error("No error from Load() with no file")
	}

	if conf.initialized {
		t.Error("conf failed to load but said it was initialized.")
	}
}

func TestConfiguration_Load_InvalidFile(t *testing.T) {
	conf := Configuration{}
	if err := conf.Load("application.invalid.yaml"); err == nil {
		t.Error("No error from Load() with invalid file")
	}

	if conf.initialized {
		t.Error("conf failed to load but said it was initialized.")
	}
}
