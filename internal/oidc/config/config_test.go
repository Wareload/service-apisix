package config

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestValidConfiguration(t *testing.T) {
	input := Configuration{
		Auth: Auth{
			ClientId:     "client-id",
			ClientSecret: "client-secret",
			Scopes:       "openid profile email",
			Leeway:       15,
		},
		UrlPaths: UrlPaths{
			BaseUrl:       "http://localhost:3000",
			LoginPath:     "/login",
			CallbackPath:  "/callback",
			PostLoginUrl:  "/post-login",
			UserinfoPath:  "/userinfo",
			LogoutPath:    "/logout",
			PostLogoutUrl: "/post-logout",
		},
		WellKnown: WellKnown{
			Issuer:                "my-issuer",
			AuthorizationEndpoint: "http://localhost:8080/auth",
			TokenEndpoint:         "http://localhost:8080/token",
			UserinfoEndpoint:      "http://localhost:8080/userinfo",
			RevocationEndpoint:    "http://localhost:8080/revoke",
		},
		Cookie: Cookie{
			Name:     "auth",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: "lax",
			Secret:   "FOhGpqouMhuYQfSySIPQEKRfdcqnwBeF",
		},
		Invalid: false,
	}
	marshal, err := yaml.Marshal(input)
	if err != nil {
		t.Error("failed to marshal configuration")
	}
	configuration := Configuration{}
	err = yaml.Unmarshal(marshal, &configuration)
	if err != nil {
		t.Errorf("failed to unmarshal configuration: %s", err)
	}
	configuration.Validate()
	if configuration.Invalid {
		t.Errorf("invalid configuration")
	}
}

func TestInvalidAuthConfiguration(t *testing.T) {
	input := Configuration{
		Auth: Auth{},
		UrlPaths: UrlPaths{
			BaseUrl:       "http://localhost:3000",
			LoginPath:     "/login",
			CallbackPath:  "/callback",
			PostLoginUrl:  "/post-login",
			UserinfoPath:  "/userinfo",
			LogoutPath:    "/logout",
			PostLogoutUrl: "/post-logout",
		},
		WellKnown: WellKnown{
			Issuer:                "my-issuer",
			AuthorizationEndpoint: "http://localhost:8080/auth",
			TokenEndpoint:         "http://localhost:8080/token",
			UserinfoEndpoint:      "http://localhost:8080/userinfo",
			RevocationEndpoint:    "http://localhost:8080/revoke",
		},
		Cookie: Cookie{
			Name:     "auth",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: "lax",
			Secret:   "FOhGpqouMhuYQfSySIPQEKRcqnwBeF",
		},
		Invalid: false,
	}
	marshal, err := yaml.Marshal(input)
	if err != nil {
		t.Error("failed to marshal configuration")
	}
	configuration := Configuration{}
	err = yaml.Unmarshal(marshal, &configuration)
	if err != nil {
		t.Errorf("failed to unmarshal configuration: %s", err)
	}
	configuration.Validate()
	if !configuration.Invalid {
		t.Errorf("invalid configuration is valid")
	}
}

func TestInvalidUrlPathsConfiguration(t *testing.T) {
	input := Configuration{
		Auth: Auth{
			ClientId:     "client-id",
			ClientSecret: "client-secret",
			Scopes:       "openid profile email",
			Leeway:       15,
		},
		UrlPaths: UrlPaths{},
		WellKnown: WellKnown{
			Issuer:                "my-issuer",
			AuthorizationEndpoint: "http://localhost:8080/auth",
			TokenEndpoint:         "http://localhost:8080/token",
			UserinfoEndpoint:      "http://localhost:8080/userinfo",
			RevocationEndpoint:    "http://localhost:8080/revoke",
		},
		Cookie: Cookie{
			Name:     "auth",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: "lax",
			Secret:   "FOhGpqouMhuYQfSySIPQEKRcqnwBeF",
		},
		Invalid: false,
	}
	marshal, err := yaml.Marshal(input)
	if err != nil {
		t.Error("failed to marshal configuration")
	}
	configuration := Configuration{}
	err = yaml.Unmarshal(marshal, &configuration)
	if err != nil {
		t.Errorf("failed to unmarshal configuration: %s", err)
	}
	configuration.Validate()
	if !configuration.Invalid {
		t.Errorf("invalid configuration is valid")
	}
}

func TestInvalidWellknownConfiguration(t *testing.T) {
	input := Configuration{
		Auth: Auth{
			ClientId:     "client-id",
			ClientSecret: "client-secret",
			Scopes:       "openid profile email",
			Leeway:       15,
		},
		UrlPaths: UrlPaths{
			BaseUrl:       "http://localhost:3000",
			LoginPath:     "/login",
			CallbackPath:  "/callback",
			PostLoginUrl:  "/post-login",
			UserinfoPath:  "/userinfo",
			LogoutPath:    "/logout",
			PostLogoutUrl: "/post-logout",
		},
		WellKnown: WellKnown{},
		Cookie: Cookie{
			Name:     "auth",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: "lax",
			Secret:   "FOhGpqouMhuYQfSySIPQEKRcqnwBeF",
		},
		Invalid: false,
	}
	marshal, err := yaml.Marshal(input)
	if err != nil {
		t.Error("failed to marshal configuration")
	}
	configuration := Configuration{}
	err = yaml.Unmarshal(marshal, &configuration)
	if err != nil {
		t.Errorf("failed to unmarshal configuration: %s", err)
	}
	configuration.Validate()
	if !configuration.Invalid {
		t.Errorf("invalid configuration is valid")
	}
}

func TestInvalidCookieConfiguration(t *testing.T) {
	input := Configuration{
		Auth: Auth{
			ClientId:     "client-id",
			ClientSecret: "client-secret",
			Scopes:       "openid profile email",
			Leeway:       15,
		},
		UrlPaths: UrlPaths{
			BaseUrl:       "http://localhost:3000",
			LoginPath:     "/login",
			CallbackPath:  "/callback",
			PostLoginUrl:  "/post-login",
			UserinfoPath:  "/userinfo",
			LogoutPath:    "/logout",
			PostLogoutUrl: "/post-logout",
		},
		WellKnown: WellKnown{
			Issuer:                "my-issuer",
			AuthorizationEndpoint: "http://localhost:8080/auth",
			TokenEndpoint:         "http://localhost:8080/token",
			UserinfoEndpoint:      "http://localhost:8080/userinfo",
			RevocationEndpoint:    "http://localhost:8080/revoke",
		},
		Cookie:  Cookie{},
		Invalid: false,
	}
	marshal, err := yaml.Marshal(input)
	if err != nil {
		t.Error("failed to marshal configuration")
	}
	configuration := Configuration{}
	err = yaml.Unmarshal(marshal, &configuration)
	if err != nil {
		t.Errorf("failed to unmarshal configuration: %s", err)
	}
	configuration.Validate()
	if !configuration.Invalid {
		t.Errorf("invalid configuration is valid")
	}
}
