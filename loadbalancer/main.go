package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

type Instance interface {
	URL() string
	Healthcheck() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type instance struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func (s *instance) URL() string { return s.addr }

func (s *instance) Healthcheck() bool {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	_, err := client.Get(s.URL() + "health")
	return err == nil
}

func (s *instance) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func newInstance(addr string) *instance {
	serverUrl, err := url.Parse(addr)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	return &instance{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port      string
	rbc       int
	instances []Instance
}

func NewLoadBalancer(port string, instances []Instance) *LoadBalancer {
	return &LoadBalancer{
		port:      port,
		rbc:       0,
		instances: instances,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Instance {
	instance := lb.instances[lb.rbc%len(lb.instances)]
	count := 0
	for !instance.Healthcheck() {
		lb.rbc++
		instance = lb.instances[lb.rbc%len(lb.instances)]
		count++
		if count >= len(lb.instances) {
			return nil
		}
	}
	lb.rbc++
	return instance
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetInstance := lb.getNextAvailableServer()
	if targetInstance == nil {
		rw.WriteHeader(500)
		return
	}
	fmt.Printf("forwarding request to URL %q\n", targetInstance.URL())
	targetInstance.Serve(rw, req)
}

func main() {
	instances := []Instance{
		newInstance("http://dis-cloud-1/"),
		newInstance("http://dis-cloud-2/"),
		newInstance("http://dis-cloud-3/"),
	}

	lb := NewLoadBalancer("8000", instances)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
