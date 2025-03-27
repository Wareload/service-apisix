package config

import "testing"

func TestValidCookie(t *testing.T) {
	cookie := Cookie{
		Name:     "auth",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: "lax",
		Secret:   "FOhGpqouMhuYQfSySIPQEKRcqnwBeF",
	}
	if !cookie.isValid() {
		t.Error("valid cookie is invalid")
	}
}

func TestInvalidSecretCookie(t *testing.T) {
	cookie := Cookie{
		Name:     "auth",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: "lax",
		Secret:   "FOhGpqouMhuYQfSySIPQEKRcqnwBe",
	}
	if cookie.isValid() {
		t.Error("invalid cookie is valid")
	}
}

func TestValidDefaultsCookie(t *testing.T) {
	cookie := Cookie{
		Path:   "/",
		Secret: "FOhGpqouMhuYQfSySIPQEKRcqnwBea",
	}
	if !cookie.isValid() {
		t.Error("valid cookie is invalid")
	}
}

func TestValidMinimalCookie(t *testing.T) {
	cookie := Cookie{
		Secret: "FOhGpqouMhuYQfSySIPQEKRcqnwBea",
	}
	if !cookie.isValid() {
		t.Error("valid cookie is invalid")
	}
}
