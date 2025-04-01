package oidc

import (
	"fmt"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/routes"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
	"net/http"
	"strings"
)

type Oidc struct {
	plugin.DefaultPlugin
}

func (s Oidc) Name() string {
	return "oidc"
}

func (s Oidc) ParseConf(in []byte) (interface{}, error) {
	// never return an error, otherwise the plugin will be disabled
	conf := config.Configuration{
		Cookie: config.Cookie{
			Secure:   true,
			HttpOnly: true,
			SameSite: "lax",
		},
		Invalid: false,
	}
	err := yaml.Unmarshal(in, &conf)
	if err != nil {
		log.Errorf("failed to unmarshal config")
		return conf, nil
	}
	conf.Validate()
	conf.OAuth = &oauth2.Config{
		ClientID:     conf.Auth.ClientId,
		ClientSecret: conf.Auth.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", conf.UrlPaths.BaseUrl, conf.UrlPaths.CallbackPath),
		Scopes:       strings.Split(conf.Auth.Scopes, " "),
		Endpoint: oauth2.Endpoint{
			AuthURL:   conf.WellKnown.AuthorizationEndpoint,
			TokenURL:  conf.WellKnown.TokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	if conf.Invalid {
		log.Errorf("config is invalid")
		return conf, nil
	}
	return s, nil
}

func (s Oidc) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	configuration, ok := conf.(config.Configuration)
	if !ok || configuration.Invalid {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	switch string(r.Path()) {
	case configuration.UrlPaths.LoginPath:
		routes.HandleLogin(configuration, w, r)
	case configuration.UrlPaths.LogoutPath:
		routes.HandleLogout(configuration, w, r)
	case configuration.UrlPaths.CallbackPath:
		routes.HandleCallback(configuration, w, r)
	case configuration.UrlPaths.UserinfoPath:
		routes.HandleUserinfo(configuration, w, r)
	default:
		routes.HandleProxy(configuration, w, r)
	}
}
