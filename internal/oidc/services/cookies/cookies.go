package cookies

import (
	"github.com/Wareload/service-apisix/internal/oidc/config"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

const maxCookieSize = 3800

type AuthFlowCookie struct {
	State string `json:"state"`
	Nonce string `json:"nonce"`
}

type AuthAccessCookie struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetAuthFlowCookie() (flow AuthFlowCookie, err error) {
	return flow, err
}

func GetAuthAccessCookie() (tokens AuthAccessCookie, err error) {
	return tokens, nil
}

func SetAuthFlowCookie(w http.ResponseWriter, config config.Configuration, authFlow AuthFlowCookie) error {
	return nil
}

func SetAuthAccessCookie(w http.ResponseWriter, config config.Configuration, accessToken, refreshToken string) error {
	return nil
}

func DeleteCookies(w http.ResponseWriter, config config.Configuration) {

}

func RemovePluginCookiesFromRequestHeader(r pkgHTTP.Request, config config.Configuration) {

}

func setChunkedCookie() {

}
func deleteChunkedCookie() {

}
