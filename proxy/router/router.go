package Router

import (
	"fmt"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
	s "github.com/lautarojayat/auth_proxy/proxy/strategy"
)


func NewRouter(rp *httputil.ReverseProxy) (*httprouter.Router){
	r:= httprouter.New()
	r.GET("/*path", s.CheckAuthHeader(rp) )
	r.POST("/*path", s.CheckAuthHeader(rp))
	r.DELETE("/*path", s.CheckAuthHeader(rp))
	r.PUT("/*path", s.CheckAuthHeader(rp))
	r.OPTIONS("/*paths", s.FwdOptionsReq(rp))
	fmt.Println("Routes Registered")
	return r
}