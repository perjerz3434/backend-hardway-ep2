package main

import (
	"bytes"
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strings"
)

func main() {
	http.ListenAndServe(":8082", http.HandlerFunc(index))
}

func index(w http.ResponseWriter, r *http.Request) {
	username, password := parseBasicAuth(r)
	if username != "root" || subtle.ConstantTimeCompare([]byte(password), []byte("toor")) != 1 {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Enter your password\"")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if r.URL.Path == "/logout" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Hello World"))
}

func parseBasicAuth(r *http.Request) (username, password string) {
	auth := r.Header.Get("Authorization")
	if len(auth) <= 6 {
		return
	}
	if !strings.EqualFold(auth[:6], "Basic ") {
		return
	}
	userpass, _ := base64.StdEncoding.DecodeString(auth[6:])
	delim := bytes.IndexByte(userpass, ':')
	if delim == -1 {
		return
	}
	return string(userpass[:delim]), string(userpass[delim+1:])
}
