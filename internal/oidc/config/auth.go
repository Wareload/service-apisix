package config

import "github.com/apache/apisix-go-plugin-runner/pkg/log"

type Auth struct {
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Scopes       string `yaml:"scopes"`
	Leeway       int    `yaml:"leeway"`
}

func (a *Auth) isValid() bool {
	if a.ClientId == "" || a.ClientSecret == "" {
		log.Errorf("missing client_id or client_secret")
		return false
	}
	if a.Scopes == "" {
		a.Scopes = "openid"
	}
	return true
}
