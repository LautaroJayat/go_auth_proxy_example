package main

import (
	"fmt"
	"log"

	c "github.com/lautarojayat/auth_proxy/proxy/config"
)

func main() {
	fmt.Println("Starting proxy server")

	configs := c.GetConfigs()

	fmt.Printf("PROXY_HOST %v\n", configs.ProxyHost)
	fmt.Printf("TARGET_HOST %v\n", configs.TargetHost)
	fmt.Printf("SCHEME %v\n", configs.Scheme)
	fmt.Printf("Auth %v\n", configs.SecretString)

	srv := createServer(configs)

	go startServer(srv)

	stop, ctx, cancel := notifyWhenDone()

	<-stop

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	fmt.Println("Proxy has received a kill/terminate signal, bye bye!")
}
