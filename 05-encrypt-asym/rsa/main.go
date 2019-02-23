package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	msg := []byte("hello, superman please save my cat")
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	x509PrivKey := x509.MarshalPKCS1PrivateKey(privKey)
	fmt.Println(base64.StdEncoding.EncodeToString(x509PrivKey))
	fmt.Println("-----------")

	x509PubKey := x509.MarshalPKCS1PublicKey(&privKey.PublicKey)
	fmt.Println(base64.StdEncoding.EncodeToString(x509PubKey))

	keyPemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509PrivKey,
	}
	pem.Encode(os.Stdout, keyPemBlock)

	fmt.Println("-----------")
	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, &privKey.PublicKey, msg)
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
	plaintext, _ := rsa.DecryptPKCS1v15(rand.Reader, privKey, ciphertext)
	fmt.Println(string(plaintext))
}
