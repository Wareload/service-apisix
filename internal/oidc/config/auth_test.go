package config

import "testing"

func TestValidAuth(t *testing.T) {
	auth := Auth{
		ClientId:     "client-id",
		ClientSecret: "client-secret",
		Scopes:       "openid profile email",
		Leeway:       15,
	}
	if !auth.isValid() {
		t.Error("valid auth is invalid")
	}
}

func TestInvalidClientIdAuth(t *testing.T) {
	auth := Auth{
		ClientSecret: "client-secret",
		Scopes:       "openid profile email",
		Leeway:       15,
	}
	if auth.isValid() {
		t.Error("invalid auth is valid")
	}
}

func TestValidClientSecretAuth(t *testing.T) {
	auth := Auth{
		ClientId: "client-id",
		Scopes:   "openid profile email",
		Leeway:   15,
	}
	if auth.isValid() {
		t.Error("invalid auth is valid")
	}
}

func TestValidScopeAuth(t *testing.T) {
	auth := Auth{
		ClientId:     "client-id",
		ClientSecret: "client-secret",
		Leeway:       15,
	}
	if !auth.isValid() {
		t.Error("valid auth is invalid")
	}
}
