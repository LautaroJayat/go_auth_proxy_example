package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"time"

	c "github.com/lautarojayat/go_auth_proxy_example/proxy/config"
	d "github.com/lautarojayat/go_auth_proxy_example/proxy/director"
	rt "github.com/lautarojayat/go_auth_proxy_example/proxy/router"
)

func createReverseProxy(configs *c.Configs) httputil.ReverseProxy {
	return httputil.ReverseProxy{
		Director: d.NewDirector(configs),
	}
}

func startServer(srv http.Server) {
	fmt.Println("starting server")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func notifyWhenDone() (chan os.Signal, context.Context, context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return stop, ctx, cancel
}

func createServer(configs *c.Configs) http.Server {
	rp := createReverseProxy(configs)
	r := rt.NewRouter(&rp, configs)
	srv := http.Server{
		Handler: r,
		Addr:    ":8000",
	}
	return srv
}
