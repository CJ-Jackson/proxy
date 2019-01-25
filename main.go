package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	fmt.Println("Proxy Running @ 8888")
	log.Fatal(http.ListenAndServe(":8888", proxy))
}
