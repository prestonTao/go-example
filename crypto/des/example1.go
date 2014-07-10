package main

import (
	"bytes"
	"crypto/des"
	"fmt"
)

func main() {
	example1()
}

func example1() {
	text := []byte("example key 1234")
	key := []byte("12345678keytaopopooooooo")

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	padding := block.BlockSize() - len(text)%block.BlockSize()
	padtext := bytes.Repeat([]byte{0}, padding)
	srcText := append(text, padtext...)

	dst := make([]byte, len(text))
	block.Encrypt(dst, text)

	fmt.Println(dst)

	//解密
	deBlock, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	deDst := make([]byte, len(text))
	deBlock.Decrypt(deDst, dst)
	fmt.Println(string(deDst))
}
