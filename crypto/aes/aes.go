package main

import (
	"crypto/aes"
	"fmt"
	"strings"
)

func main() {
	//////////////------AES加密------//////////////
	//秘钥 16/24/32bytes对应AES-128/AES-192/AES-256.
	key := []byte{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 0, 1, 2, 3, 4, 5, 6,
		7, 8, 9, 0, 1, 2, 3, 4,
		5, 6, 7, 8, 9, 0, 1, 2,
	}
	fmt.Println("每次加密的字节数：", aes.BlockSize)
	//明文
	cleartext := make([]byte, aes.BlockSize)
	strings.NewReader("I'm a cleartext!").Read(cleartext)
	//密文
	ciphertext := make([]byte, aes.BlockSize)
	cip, _ := aes.NewCipher(key)
	//加密
	cip.Encrypt(ciphertext, cleartext)
	fmt.Println("明文：", cleartext)
	fmt.Println("密文：", ciphertext)
	//解密
	cip.Decrypt(cleartext, ciphertext)
	fmt.Println("密文：", ciphertext)
	fmt.Println("明文：", cleartext)
	fmt.Printf("明文： %s", cleartext)
	//////////////------AES加密------//////////////
}
