package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(proxy))

}

func proxy(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.RequestURI)
	if r.Method == http.MethodConnect {
		// HTTPS
		// CONNECT www.google.com
		upstreamConn, err := net.Dial("tcp", r.RequestURI)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		defer upstreamConn.Close()

		hijacker, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "Client not support hijacker", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		respConn, _, err := hijacker.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// End Connect

		// TLS Handshake
		// GET / HTTP/1.1

		// read conn, write respConn
		go io.Copy(upstreamConn, respConn)
		// read respConn
		io.Copy(respConn, upstreamConn)
		return
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()
}
