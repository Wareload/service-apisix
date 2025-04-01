package routes

import (
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/services/cookies"
	"github.com/Wareload/service-apisix/internal/oidc/services/oidc"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

func HandleUserinfo(config config.Configuration, w http.ResponseWriter, r pkgHTTP.Request) {
	if r.Method() != "GET" {
		onMethodNotAllowed(w)
		return
	}
	tokens, err := cookies.GetAuthAccessCookie(r, config)
	if err != nil {
		onUnauthorized(r, w, config)
		return
	}
	currentAccessToken, refreshed, err := updateTokensIfNeeded(r, w, config, tokens.AccessToken, tokens.RefreshToken, r.Context())
	if err != nil {
		onUnauthorized(r, w, config)
		return
	}
	if refreshed {
		onTemporaryRedirect(w, r)
		return
	}
	response, err := oidc.GetUserInfo(currentAccessToken, config.WellKnown.UserinfoEndpoint)
	if err != nil {
		onInternalServerError(w, err)
		return
	}
	_, err = w.Write([]byte(response))
	if err != nil {
		onInternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
