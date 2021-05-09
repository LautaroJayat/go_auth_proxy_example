package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	d "github.com/lautarojayat/auth_proxy/proxy/director"
	rt "github.com/lautarojayat/auth_proxy/proxy/router"
)

var rp = httputil.ReverseProxy{
	Director: d.NewDirector(),
}

func main(){

	fmt.Printf("SCHEME %v\n", os.Getenv("SCHEME"))
	fmt.Printf("HOST %v\n", os.Getenv("HOST"))
	fmt.Printf("Auth %v\n", os.Getenv("AUTH"))
	fmt.Println("Starting proxy server")	
	
	r := rt.NewRouter(&rp)

	srv := http.Server{
		Handler: r,
		Addr: "localhost:8081",
	}
	srv.ListenAndServe()
}