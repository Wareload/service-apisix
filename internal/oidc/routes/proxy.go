package routes

import (
	"github.com/Wareload/service-apisix/internal/oidc/config"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"net/http"
)

func HandleProxy(config config.Configuration, w http.ResponseWriter, r pkgHTTP.Request) {

}
