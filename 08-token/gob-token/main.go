package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"
	"time"
)

var secret = []byte("supersecret")

func main() {
	token := generateToken(Token{
		UserID: "123",
	}, secret)
	fmt.Println(token)

	t := parseToken(token, secret)
	fmt.Print(t)
}

type Token struct {
	UserID  string
	IssueAt int64
}

//GOBเป็น JSON ที่อ่านได้เฉพาะ Go -> JSON, YAML, TOML, PEM
//RS256 RSA SHA256
//ES512 -> ECDSA sSHA256
//base64(gob(data)).base64(signature)
//RSA PSS
//HS256 HMAC-SHA256 Symmatric
//
func generateToken(t Token, key []byte) string {
	t.IssueAt = time.Now().Unix()

	buf := bytes.Buffer{}
	gob.NewEncoder(&buf).Encode(t) // <-- encode ใส่ตัวแปร buff JSON.stringrify

	h := hmac.New(sha256.New, key)
	h.Write(buf.Bytes())
	signature := h.Sum(nil)

	return encodeBase64(buf.Bytes()) + "." + encodeBase64(signature)
}

func encodeBase64(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	p, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return p
}

func parseToken(token string, key []byte) *Token {
	//aaaa.bbb
	ts := strings.Split(token, ".")
	if len(ts) != 2 {
		return nil
	}
	data := decodeBase64(ts[0])
	signature := decodeBase64(ts[1])

	h := hmac.New(sha256.New, key)
	h.Write(data)
	dataSignature := h.Sum(nil)
	if !hmac.Equal(signature, dataSignature) {
		return nil
	}
	// signature valid
	var t Token
	gob.NewDecoder(bytes.NewReader(data)).Decode(&t)
	return &t
}
