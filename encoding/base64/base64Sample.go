package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	example2()
}

func example2() {
	src := "nihaonihao"
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	fmt.Println([]byte(src))
	dst := b64.EncodeToString([]byte(src))
	fmt.Println(dst)

	srcToo, _ := b64.DecodeString(dst)
	fmt.Println(srcToo)
}

func example() {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	println(b64.EncodeToString([]byte("哈哈")))
	s := []byte("http://golang.org/pkg/encoding/base64/#variables")
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(s))
}
