package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.Host != "localhost:8081" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if os.Getenv("SECRET") != "" && r.Header.Get("SECRET") != os.Getenv("SECRET") {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.Header().Add("x-method", r.Method)
	w.Header().Add("x-url", r.RequestURI)
	w.Header().Add("x-host", r.Host)
    io.WriteString(w, "hi from target server!\n")
}


func main(){
	fmt.Println("Starting target server")

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcome)


	srv := http.Server{
		Handler: mux,
		Addr: "localhost:8080",
	}
	srv.ListenAndServe()
}