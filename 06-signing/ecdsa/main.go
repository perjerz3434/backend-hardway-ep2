package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

// ECDSA-SHA256
func main() {
	msg := []byte("heello, superman")

	digest := sha256.Sum256(msg)
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//signature, _ := privKey.Sign(rand.Reader, digest[:], nil)
	//fmt.Println(base64.StdEncoding.EncodeToString(signature))
	r, s, _ := ecdsa.Sign(rand.Reader, privKey, digest[:])

	x509Priv, _ := x509.MarshalECPrivateKey(privKey)
	pem.Encode(os.Stdout, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: x509Priv,
	})

	x509Pub, _ := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	pem.Encode(os.Stdout, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: x509Pub,
	})

	signature := append(r.Bytes(), s.Bytes()...)
	// Send base 64 to another person
	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	// Verify
	r = new(big.Int).SetBytes(signature[:32])
	s = new(big.Int).SetBytes(signature[32:])

	ok := ecdsa.Verify(&privKey.PublicKey, digest[:], r, s)
	fmt.Println(ok)
}
