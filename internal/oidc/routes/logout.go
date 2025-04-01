package routes

import (
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/services/cookies"
	"github.com/Wareload/service-apisix/internal/oidc/services/oidc"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

func HandleLogout(config config.Configuration, w http.ResponseWriter, r pkgHTTP.Request) {
	if r.Method() != "GET" && r.Method() != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tokens, err := cookies.GetAuthAccessCookie()
	cookies.DeleteCookies(w, config)
	if err != nil {
		onRedirect(w, config.UrlPaths.PostLogoutUrl)
		return
	}
	_ = oidc.RevokeTokens(tokens.RefreshToken, config)
	onRedirect(w, config.UrlPaths.PostLoginUrl)
}
