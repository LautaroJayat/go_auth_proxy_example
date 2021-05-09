package strategy

import (
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/julienschmidt/httprouter"
)

func CheckAuthHeader(rp *httputil.ReverseProxy) httprouter.Handle{
	//Getting envs
	secret := os.Getenv("SECRET")
	hardCodedAuth := os.Getenv("AUTH")
	//Returning our handler
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

		if (secret!= ""){
			r.Header.Add("SECRET", secret)
		}

		auth := r.Header.Get("Auth")

		if auth != hardCodedAuth {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "Auth needed\n")
			return
		}

		rp.ServeHTTP(w, r)
	}
}
