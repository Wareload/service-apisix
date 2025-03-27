package config

import "testing"

func TestInvalidMinimalUrlPaths(t *testing.T) {
	urlPaths := UrlPaths{}
	if urlPaths.isValid() {
		t.Errorf("urlPaths should be invalid")
	}
}

func TestValidUrlPaths(t *testing.T) {
	urlPaths := UrlPaths{
		BaseUrl:      "http://localhost:7070",
		LoginPath:    "/login",
		LogoutPath:   "/logout",
		CallbackPath: "/callback",
		UserinfoPath: "/userinfo",
	}
	if !urlPaths.isValid() {
		t.Errorf("urlPaths should be valid")
	}
}
func TestInvalidUrlPaths(t *testing.T) {
	urlPaths := UrlPaths{
		LoginPath:    "/login",
		LogoutPath:   "/logout",
		CallbackPath: "/callback",
		UserinfoPath: "/userinfo",
	}
	if urlPaths.isValid() {
		t.Errorf("urlPaths should be valid")
	}
}

func TestValidMinimalUrlPaths(t *testing.T) {
	urlPaths := UrlPaths{
		BaseUrl: "http://localhost:7070",
	}
	if !urlPaths.isValid() {
		t.Errorf("urlPaths should be valid")
	}
}
