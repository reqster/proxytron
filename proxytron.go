package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// a reliably HTTP-accepting website
	target, _ := url.Parse("http://httpforever.com")
	proxy := httputil.NewSingleHostReverseProxy(target)
	http.ListenAndServe(":8068", proxy)
}