package oidc

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"golang.org/x/oauth2"
)

func GenerateLoginUrl(config config.Configuration) (string, string, string, error) {
	state, errState := randomString()
	nonce, errNonce := randomString()
	loginUrl := config.OAuth.AuthCodeURL(state, oauth2.SetAuthURLParam("nonce", nonce))
	return loginUrl, state, nonce, errors.Join(errState, errNonce)
}

func randomString() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
