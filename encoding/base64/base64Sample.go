package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	println(b64.EncodeToString([]byte("哈哈")))
	s := []byte("http://golang.org/pkg/encoding/base64/#variables")
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(s))
}
