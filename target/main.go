package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting target server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", createWelcomeController(os.Getenv("SECRET")))

	srv := http.Server{
		Handler: mux,
		Addr:    "localhost:8080",
	}
	srv.ListenAndServe()
}
