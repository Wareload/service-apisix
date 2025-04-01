package oidc

import (
	"context"
	"fmt"
	"github.com/Wareload/service-apisix/internal/oidc/config"
)

func ExchangeCodeToTokens(config config.Configuration, code string, ctx context.Context) (string, string, string, error) {
	exchange, err := config.OAuth.Exchange(ctx, code)
	if err != nil {
		return "", "", "", err
	}
	idToken, ok := exchange.Extra("id_token").(string)
	if !ok {
		return "", "", "", fmt.Errorf("ID token not found or could not be parsed")
	}
	return exchange.AccessToken, idToken, exchange.RefreshToken, nil
}
