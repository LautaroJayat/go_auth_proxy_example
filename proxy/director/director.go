package director

import (
	"net/http"

	"github.com/lautarojayat/auth_proxy/proxy/config"
)

func NewDirector(config *config.Configs) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Host = config.TargetHost
		req.URL.Scheme = config.Scheme
		req.Header.Add("x-secret", config.TargetSecret)
	}
}
