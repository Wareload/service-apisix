package config

import (
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"net/http"
)

type Cookie struct {
	Name     string `yaml:"name"`
	Path     string `yaml:"path"`
	Secure   bool   `yaml:"secure"`
	HttpOnly bool   `yaml:"http_only"`
	SameSite string `yaml:"same_site"`
	Secret   string `yaml:"secret"`
}

func (c *Cookie) isValid() bool {
	if c.Name == "" {
		c.Name = "auth"
	}
	if c.Path == "" {
		c.Path = "/"
	}
	if !(c.SameSite != "lax" && c.SameSite != "strict" && c.SameSite != "none") {
		c.SameSite = "lax"
	}
	if len(c.Secret) != 30 {
		log.Errorf("cookie secret must be 30 chars long")
		return false
	}
	return true
}

func (c *Cookie) GetCookieSameSite() http.SameSite {
	switch c.SameSite {
	case "lax":
		return http.SameSiteLaxMode
	case "strict":
		return http.SameSiteStrictMode
	case "none":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}
