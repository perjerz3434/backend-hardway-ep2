package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	msg := []byte("hello, superman please save my cat")
	key := make([]byte, 32)
	{
		buf := bytes.NewBuffer(key)
		buf.Reset()
		buf.WriteString("supersecret")
	}
	fmt.Println(string(key))
	// encrypt bl
	block, _ := aes.NewCipher(key)
	// Create แม่พิมพ์
	aesgcm, _ := cipher.NewGCM(block)
	//put block in box and seal
	// gen nonce
	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce)
	// เอาแม่พิมพ์มาสร้าง
	ciphertext := aesgcm.Seal(nil, nonce, msg, nil)
	encodedCipherText := base64.RawStdEncoding.EncodeToString(ciphertext)
	fmt.Println(encodedCipherText)
	//need to attached nonce to open box

	plaintext, _ := aesgcm.Open(nil, nonce, ciphertext, nil)
	fmt.Println(string(plaintext))
}
