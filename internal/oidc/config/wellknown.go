package config

import "github.com/apache/apisix-go-plugin-runner/pkg/log"

type WellKnown struct {
	Issuer                string `yaml:"issuer"`
	AuthorizationEndpoint string `yaml:"authorization_endpoint"`
	TokenEndpoint         string `yaml:"token_endpoint"`
	UserinfoEndpoint      string `yaml:"userinfo_endpoint"`
	RevocationEndpoint    string `yaml:"revocation_endpoint"`
}

func (w *WellKnown) isValid() bool {
	if w.Issuer != "" && w.AuthorizationEndpoint != "" && w.TokenEndpoint != "" && w.UserinfoEndpoint != "" && w.RevocationEndpoint != "" {
		return true
	}
	log.Errorf("wellknown configuration is incomplete")
	return false
}
