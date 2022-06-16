package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func createWelcomeController(secret string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		if r.Host != os.Getenv("PROXY_HOST") {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "The request didn't came from the proxy\n")
			return
		}
		if r.Header.Get("x-secret") != secret {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "The didn't include the correct secret string in 'x-secret' header\n")
			return
		}
		w.Header().Add("x-method", r.Method)
		w.Header().Add("x-url", r.RequestURI)
		w.Header().Add("x-host", r.Host)
		io.WriteString(w, "Hi from the target server!\n")
	}
}
