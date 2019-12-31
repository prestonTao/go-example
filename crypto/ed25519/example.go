package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/ed25519"
)

func main() {
	puk, prk, _ := ed25519.GenerateKey(rand.Reader)

	//签名文件
	text := []byte("test message")
	//生成签名
	sig := ed25519.Sign(prk, text)
	//验证签名
	ok := ed25519.Verify(puk, text, sig)

	fmt.Println(ok)

}
