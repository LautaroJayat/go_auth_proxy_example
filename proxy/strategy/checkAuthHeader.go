package strategy

import (
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/julienschmidt/httprouter"
)

func MakeCheckSecretInAuthHeader(rp *httputil.ReverseProxy, hardCodedAuth string) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		auth := r.Header.Get("Authorization")
		if auth != hardCodedAuth {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "Auth needed\n")
			return
		}
		rp.ServeHTTP(w, r)
	}
}
