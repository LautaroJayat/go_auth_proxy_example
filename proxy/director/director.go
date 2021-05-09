package director

import (
	"net/http"
	"os"
)


func NewDirector() func(req *http.Request){
	host:= os.Getenv("HOST")
	scheme:= os.Getenv("SCHEME")
	return func(req *http.Request){
		req.URL.Host = host
		req.URL.Scheme = scheme
	}
}