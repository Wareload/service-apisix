package oidc

import (
	"fmt"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"net/http"
	"net/url"
	"strings"
)

func RevokeTokens(refreshToken string, conf config.Configuration) error {
	data := url.Values{}
	data.Set("client_id", conf.Auth.ClientId)
	data.Set("client_secret", conf.Auth.ClientSecret)
	data.Set("token", refreshToken)
	req, err := http.NewRequest("POST", conf.WellKnown.RevocationEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to logout: %s", resp.Status)
	}
	return nil
}
