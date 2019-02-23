package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func main() {
	var input string
	fmt.Print("Input: ")
	fmt.Scanf("%s", &input)
	output := hash1(input)
	fmt.Println("Output:", output)
}

func hash1(s string) []byte {
	b := sha256.Sum256([]byte(s))
	return b[:]
}

func hash2(s string) []byte {
	h := sha512.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func encode(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}
