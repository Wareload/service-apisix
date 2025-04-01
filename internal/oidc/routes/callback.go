package routes

import (
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/services/cookies"
	"github.com/Wareload/service-apisix/internal/oidc/services/oidc"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

func HandleCallback(config config.Configuration, w http.ResponseWriter, r pkgHTTP.Request) {
	if r.Method() != "GET" {
		onMethodNotAllowed(w)
		return
	}
	queryParams := r.Args()
	state := queryParams.Get("state")
	iss := queryParams.Get("iss")
	code := queryParams.Get("code")
	flow, err := cookies.GetAuthFlowCookie()
	if err != nil {
		onUnauthorized(w, config)
		return
	}
	if state != flow.State || iss != config.WellKnown.Issuer {
		onUnauthorized(w, config)
		return
	}
	accessToken, idToken, refreshToken, err := oidc.ExchangeCodeToTokens(config, code, r.Context())
	if err != nil {
		onInternalServerError(w, err)
		return
	}
	if !isNonceMatching(idToken, flow.Nonce) {
		onUnauthorized(w, config)
		return
	}
	err = cookies.SetAuthAccessCookie(w, config, accessToken, refreshToken)
	if err != nil {
		onInternalServerError(w, err)
		return
	}
	onRedirect(w, config.UrlPaths.PostLoginUrl)
}
