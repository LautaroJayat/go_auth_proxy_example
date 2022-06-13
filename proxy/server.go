package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"time"

	c "github.com/lautarojayat/auth_proxy/proxy/config"
	d "github.com/lautarojayat/auth_proxy/proxy/director"
	rt "github.com/lautarojayat/auth_proxy/proxy/router"
)

func createReverseProxy(configs *c.Configs) httputil.ReverseProxy {
	return httputil.ReverseProxy{
		Director: d.NewDirector(configs),
	}
}

func startServer(srv http.Server) {
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
		Addr:    configs.ProxyHost,
	}
	return srv
}
