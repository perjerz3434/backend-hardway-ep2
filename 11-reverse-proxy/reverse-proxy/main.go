package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	go http.ListenAndServe("9000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Println(ip)
		//proto := r.Header.Get("X-Forwarded-Proto")
		//proto := "http"
		//if r.TLS != nil {
		//	proto = "https"
		//}
		// Modify Upstream
		w.Write([]byte("Upstream server"))

	}))

	http.ListenAndServe(":8000", http.HandlerFunc(reverseProxy))
}

func reverseProxy(w http.ResponseWriter, r *http.Request) {
	//httputil.ReverseProxy{}
	r.URL.Host = "127.0.0.1:9000"
	r.URL.Scheme = "http"
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
}
