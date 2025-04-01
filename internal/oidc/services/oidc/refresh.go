package oidc

import (
	"context"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"golang.org/x/oauth2"
)

func RefreshTokens(refreshToken string, conf config.Configuration, ctx context.Context) (string, string, error) {
	token, err := conf.OAuth.TokenSource(ctx, &oauth2.Token{
		RefreshToken: refreshToken,
	}).Token()
	if err != nil {
		return "", "", err
	}
	return token.AccessToken, token.RefreshToken, nil
}
