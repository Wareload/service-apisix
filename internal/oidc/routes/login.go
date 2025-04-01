package routes

import (
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/services/cookies"
	"github.com/Wareload/service-apisix/internal/oidc/services/oidc"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

func HandleLogin(config config.Configuration, w http.ResponseWriter, r pkgHTTP.Request) {
	cookies.DeleteCookies(w, config)
	if r.Method() != "GET" && r.Method() != "POST" {
		onMethodNotAllowed(w)
		return
	}
	loginUrl, state, nonce, err := oidc.GenerateLoginUrl(config)
	if err != nil {
		onInternalServerError(w, err)
		return
	}
	err = cookies.SetAuthFlowCookie(w, config, cookies.AuthFlowCookie{
		State: state,
		Nonce: nonce,
	})
	if err != nil {
		onInternalServerError(w, err)
		return
	}
	onRedirect(w, loginUrl)
}
