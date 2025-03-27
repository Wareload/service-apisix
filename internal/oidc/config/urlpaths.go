package config

import "github.com/apache/apisix-go-plugin-runner/pkg/log"

type UrlPaths struct {
	BaseUrl       string `yaml:"base_url"`
	LoginPath     string `yaml:"login_path"`
	CallbackPath  string `yaml:"callback_path"`
	PostLoginUrl  string `yaml:"post_login_url"`
	UserinfoPath  string `yaml:"userinfo_path"`
	LogoutPath    string `yaml:"logout_path"`
	PostLogoutUrl string `yaml:"post_logout_url"`
}

func (u *UrlPaths) isValid() bool {
	if u.BaseUrl == "" {
		log.Errorf("missing base url")
		return false
	}
	if u.LoginPath == "" {
		u.LoginPath = "/login"
	}
	if u.CallbackPath == "" {
		u.CallbackPath = "/callback"
	}
	if u.PostLoginUrl == "" {
		u.PostLoginUrl = "/"
	}
	if u.UserinfoPath == "" {
		u.UserinfoPath = "/userinfo"
	}
	if u.LogoutPath == "" {
		u.LogoutPath = "/logout"
	}
	if u.PostLogoutUrl == "" {
		u.PostLogoutUrl = "/"
	}
	return true
}
