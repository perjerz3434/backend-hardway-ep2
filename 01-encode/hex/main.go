package main

import "fmt"

func main() {
	text := "hello world"
	r := hex([]byte(text))
	fmt.Println(r)
	fmt.Print(string(decode(r)))
}

const hexChartSet = "0123456789abcdef"

func hex(src []byte) string {
	r := ""
	for _, b := range src {
		r += string(hexChartSet[b>>4])
		r += string(hexChartSet[b&0xf])
	}

	return r
}

func decode(src string) []byte {
	var dst []byte
	for i := 0; i < len(src); i += 2 {
		p0 := hexToByte(src[i])   // '6' => 0x36
		p1 := hexToByte(src[i+1]) // '8' => 0x38
		b := p0<<4 | p1
		dst = append(dst, b)
	}
	return dst
}

func hexToByte(hex byte) byte {
	if hex <= 0x39 {
		return hex - 0x30
	}
	return hex - 0x61 + 0xa
}

// hex1
//      hex2
// 1111 1111
