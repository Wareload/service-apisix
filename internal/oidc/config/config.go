package config

import (
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"golang.org/x/oauth2"
)

type Configuration struct {
	Auth      Auth      `yaml:"auth"`
	UrlPaths  UrlPaths  `yaml:"url_paths"`
	WellKnown WellKnown `yaml:"well_known"`
	Cookie    Cookie    `yaml:"cookie"`
	Features  Features  `yaml:"features"`
	Invalid   bool
	OAuth     *oauth2.Config
}

func (c *Configuration) Validate() {
	c.Invalid = !(c.Auth.isValid() && c.UrlPaths.isValid() && c.WellKnown.isValid() && c.Cookie.isValid())
	if c.Invalid {
		log.Errorf("config from plugin 'oidc' is invalid")
	}
}
