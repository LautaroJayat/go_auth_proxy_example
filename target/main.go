package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting target server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", createWelcomeController(os.Getenv("TARGET_SECRET")))

	srv := http.Server{
		Handler: mux,
		Addr:    ":8000",
	}
	srv.ListenAndServe()
}
