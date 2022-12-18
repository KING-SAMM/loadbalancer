package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

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

type LoadBalancer struct {
	port			string
	roundRobinCount	int
	servers 		[]Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port: port,
		roundRobinCount: 0,
		servers: servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v \n", err)
		os.Exit(1)
	}
}

// simpleServer methods
func (s *simpleServer) Address() string { return s.addr }
func (s *simpleServer) IsAlive() bool { return bool }
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

// LoadBalaancer methods
func (lb *LoadBalancer) getNextAvailableServer() *Server {

}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {}

func main() {
	servers := []Server{
		newSimpleServer("https://favebook.com"),
		newSimpleServer("https://bing.com"),
		newSimpleServer("https://ducjducjgo.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving requests at 'localhost %s'\n", lb.port)
	http.ListenAndServe(":" +lb.port, nil)
}