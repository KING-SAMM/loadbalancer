package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

type simpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)

	handleErr(err)

	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),	
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
}