package cookies

const maxCookieSize = 3800

type AuthFlowCookie struct {
	State string `json:"state"`
	Nonce string `json:"nonce"`
}

type AuthAccessCookie struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GetAuthFlowCookie() {

}
func GetAuthAccessCookie() {

}
func DeleteCookies() {

}

func RemovePluginCookiesFromRequestHeader() {

}

func setChunkedCookie() {

}
func deleteChunkedCookie() {

}
