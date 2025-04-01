package routes

import (
	"context"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/oidc/services/cookies"
	"github.com/Wareload/service-apisix/internal/oidc/services/oidc"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func updateTokensIfNeeded(w http.ResponseWriter, conf config.Configuration, accessToken, refreshToken string, ctx context.Context) (currentAccessToken string, refreshed bool, err error) {
	accExpired, err := isTokenExpired(accessToken, conf.Auth.Leeway)
	if err != nil {
		return "", false, err
	}
	if !accExpired {
		return accessToken, false, nil
	}
	currentToken, refreshToken, err := oidc.RefreshTokens(refreshToken, conf, ctx)
	if err != nil {
		return "", false, err
	}
	return currentToken, true, cookies.SetAuthAccessCookie(w, conf, currentToken, refreshToken)
}

func isTokenExpired(raw string, leeway int) (bool, error) {
	token, _, err := jwt.NewParser().ParseUnverified(raw, jwt.MapClaims{})
	if err != nil {
		return false, err
	}
	date, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false, err
	}
	return date.Before(time.Now().Add(-time.Second * time.Duration(leeway))), err
}

func isNonceMatching(idToken string, nonce string) bool {
	token, _, err := jwt.NewParser().ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	nonceClaim, ok := claims["nonce"].(string)
	if !ok {
		return false
	}
	return nonceClaim == nonce
}

// status code helpers

func onUnauthorized(w http.ResponseWriter, conf config.Configuration) {
	cookies.DeleteCookies(w, conf)
	w.WriteHeader(http.StatusUnauthorized)
}

func onInternalServerError(w http.ResponseWriter, err error) {
	log.Errorf("Internal server error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
}

func onMethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func onRedirect(w http.ResponseWriter, redirectUrl string) {
	w.Header().Set("Location", redirectUrl)
	w.WriteHeader(http.StatusFound)
}

func onTemporaryRedirect(w http.ResponseWriter, r pkgHTTP.Request) {
	redirectURL := string(r.Path())
	query := r.Args()
	if query.Encode() != "" {
		redirectURL += "?" + query.Encode()
	}
	w.Header().Set("Location", redirectURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
