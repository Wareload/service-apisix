package routes

import (
	"fmt"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/services/cookies"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

func HandleProxy(config config.Configuration, w http.ResponseWriter, r pkgHTTP.Request) {
	if r.Header().Get("Authorization") != "" && config.Features.ByPassWithAuthHeader {
		return
	} //skip requests with authorization header
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
	cookies.RemovePluginCookiesFromRequestHeader(r, config)
	r.Header().Set("Authorization", fmt.Sprintf("Bearer %s", currentAccessToken))
}
