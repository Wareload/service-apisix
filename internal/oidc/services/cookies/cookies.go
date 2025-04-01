package cookies

import (
	"encoding/json"
	"fmt"
	"github.com/Wareload/service-apisix/internal/oidc/config"
	"github.com/Wareload/service-apisix/internal/shared/services/crypto"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
	"net/url"
	"strings"
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

func GetAuthFlowCookie(r pkgHTTP.Request, config config.Configuration) (flow AuthFlowCookie, err error) {
	raw, err := getChunkedCookie(r, config)
	if err != nil {
		return flow, err
	}
	err = json.Unmarshal(raw, &flow)
	return flow, err
}

func GetAuthAccessCookie(r pkgHTTP.Request, config config.Configuration) (tokens AuthAccessCookie, err error) {
	raw, err := getChunkedCookie(r, config)
	if err != nil {
		return tokens, err
	}
	err = json.Unmarshal(raw, &tokens)
	return tokens, err
}

func SetAuthFlowCookie(r pkgHTTP.Request, w http.ResponseWriter, config config.Configuration, authFlow AuthFlowCookie) error {
	marshal, err := json.Marshal(authFlow)
	if err != nil {
		return err
	}
	return setChunkedCookie(marshal, r, w, config)
}

func SetAuthAccessCookie(r pkgHTTP.Request, w http.ResponseWriter, config config.Configuration, accessToken, refreshToken string) error {
	marshal, err := json.Marshal(AuthAccessCookie{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		return err
	}
	return setChunkedCookie(marshal, r, w, config)
}

func DeleteCookies(r pkgHTTP.Request, w http.ResponseWriter, config config.Configuration) {
	cookies := strings.Split(r.Header().Get("Cookie"), ";")
	possibleCookies := getPossibleCookies(config)
	for _, cookie := range cookies {
		cookie = strings.TrimSpace(cookie)
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) < 2 {
			continue
		}
		cookieName := parts[0]
		_, exists := possibleCookies[cookieName]
		if exists {
			deleteCookie(w, config, cookieName)
		}
	}
}

func RemovePluginCookiesFromRequestHeader(r pkgHTTP.Request, config config.Configuration) {
	cookies := strings.Split(r.Header().Get("Cookie"), ";")
	possibleCookies := getPossibleCookies(config)
	var resultCookies []string
	for _, cookie := range cookies {
		cookie = strings.TrimSpace(cookie)
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) < 2 {
			continue
		}
		cookieName := parts[0]
		_, exists := possibleCookies[cookieName]
		if !exists {
			resultCookies = append(resultCookies, cookie)
		}
	}
	r.Header().Set("Cookie", strings.Join(resultCookies, "; "))
}

func getChunkedCookie(r pkgHTTP.Request, config config.Configuration) ([]byte, error) {
	cookies := strings.Split(r.Header().Get("Cookie"), ";")
	possibleCookies := getPossibleCookies(config)
	chunkMap := make(map[int]string)
	for _, cookie := range cookies {
		cookie = strings.TrimSpace(cookie)
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) != 2 {
			continue
		}
		name, value := parts[0], parts[1]
		index, exists := possibleCookies[name]
		if exists {
			chunkMap[index] = value
		}
	}
	var rawChunks []string
	for i := 0; i <= 10; i++ {
		if chunk, exists := chunkMap[i]; exists {
			rawChunks = append(rawChunks, chunk)
		}
	}
	raw := strings.Join(rawChunks, "")
	value, err := url.QueryUnescape(raw)
	if err != nil {
		return nil, err
	}
	decrypt, err := crypto.DecryptAES([]byte(value), []byte(config.Cookie.Secret))
	return decrypt, err
}

func setChunkedCookie(data []byte, r pkgHTTP.Request, w http.ResponseWriter, config config.Configuration) error {
	encrypted, err := crypto.EncryptAES(data, []byte(config.Cookie.Secret))
	if err != nil {
		return err
	}
	encoded := url.QueryEscape(string(encrypted))
	numChunks := (len(encoded) + maxCookieSize - 1) / maxCookieSize
	for i := 0; i < numChunks; i++ {
		start := i * maxCookieSize
		end := start + maxCookieSize
		if end > len(encoded) {
			end = len(encoded)
		}
		chunkName := fmt.Sprintf("%s%d", config.Cookie.Name, i)
		chunkValue := encoded[start:end]
		maxAge := 0
		if config.Cookie.Expiring {
			maxAge = 7 * 24 * 60 * 60
		}
		http.SetCookie(w, &http.Cookie{
			Name:     chunkName,
			Value:    chunkValue,
			HttpOnly: config.Cookie.HttpOnly,
			Secure:   config.Cookie.Secure,
			Path:     config.Cookie.Path,
			SameSite: config.Cookie.GetCookieSameSite(),
			MaxAge:   maxAge,
		})
	}
	for i := numChunks; i <= 10; i++ {
		deleteCookie(w, config, fmt.Sprintf("%s%d", config.Cookie.Name, i))
	}
	return nil
}

func getPossibleCookies(config config.Configuration) map[string]int {
	result := make(map[string]int)
	for i := 0; i <= 10; i++ {
		result[fmt.Sprintf("%s%d", config.Cookie.Name, i)] = i
	}
	return result
}

func deleteCookie(w http.ResponseWriter, config config.Configuration, cookie string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookie,
		Value:    "",
		HttpOnly: config.Cookie.HttpOnly,
		Secure:   config.Cookie.Secure,
		Path:     config.Cookie.Path,
		SameSite: config.Cookie.GetCookieSameSite(),
		MaxAge:   -1,
	})
}
