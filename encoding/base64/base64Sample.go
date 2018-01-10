package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	example3()
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

func example3() {
	str := "588583325107fa3878271759"
	// b64 := base64.NewEncoding(base64.URLEncoding)
	bs, _ := hex.DecodeString(str)
	out := base64.URLEncoding.EncodeToString(bs)
	fmt.Println(out)

	bs, _ = base64.URLEncoding.DecodeString(out)
	fmt.Println(hex.EncodeToString(bs))
}
