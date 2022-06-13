package Router

import (
	"fmt"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
	c "github.com/lautarojayat/auth_proxy/proxy/config"
	s "github.com/lautarojayat/auth_proxy/proxy/strategy"
)

func NewRouter(rp *httputil.ReverseProxy, configs *c.Configs) *httprouter.Router {
	r := httprouter.New()
	r.GET("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.SecretString))
	r.POST("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.SecretString))
	r.DELETE("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.SecretString))
	r.PUT("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.SecretString))
	r.OPTIONS("/*paths", s.FwdOptionsReq(rp))
	fmt.Println("Routes Registered")
	return r
}
