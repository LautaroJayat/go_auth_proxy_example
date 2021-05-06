package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/julienschmidt/httprouter"
)

var reverseProxy = httputil.ReverseProxy{
	Director: func(req *http.Request){
		req.URL.Host = os.Getenv("HOST")
		req.URL.Scheme = os.Getenv("SCHEME")
	},
}

func checkAuthHeader(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	auth := r.Header.Get("Auth")
	if auth == ""{
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "Auth needed\n")
		return
	}
	reverseProxy.ServeHTTP(w, r)
}


func registerRoutes(r *httprouter.Router) (*httprouter.Router){
	r.GET("/*path",checkAuthHeader )
	r.POST("/*path", checkAuthHeader)
	r.DELETE("/*path", checkAuthHeader)
	r.PUT("/*path", checkAuthHeader)
	return r
}

func main(){

	fmt.Printf("SCHEME %v\n", os.Getenv("SCHEME"))
	fmt.Printf("HOST %v\n", os.Getenv("HOST"))
	fmt.Printf("Auth %v\n", os.Getenv("Auth"))


	fmt.Println("Starting proxy server")	
	r := httprouter.New()
	r = registerRoutes(r)

	fmt.Println("Routes Registered")

	srv := http.Server{
		Handler: r,
		Addr: "localhost:8081",
	}
	srv.ListenAndServe()
}