package strategy

import (
	"net/http"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
)

func FwdOptionsReq(rp *httputil.ReverseProxy) httprouter.Handle{
	return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		rp.ServeHTTP(w, r)
	}
}