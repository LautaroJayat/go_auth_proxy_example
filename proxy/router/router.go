package Router

import (
	"fmt"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
	c "github.com/lautarojayat/go_auth_proxy_example/proxy/config"
	s "github.com/lautarojayat/go_auth_proxy_example/proxy/strategy"
)

func NewRouter(rp *httputil.ReverseProxy, configs *c.Configs) *httprouter.Router {
	r := httprouter.New()
	r.GET("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.ProxySecret))
	r.POST("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.ProxySecret))
	r.DELETE("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.ProxySecret))
	r.PUT("/*path", s.MakeCheckSecretInAuthHeader(rp, configs.ProxySecret))
	r.OPTIONS("/*path", s.FwdOptionsReq(rp))
	fmt.Println("Routes Registered")
	return r
}
