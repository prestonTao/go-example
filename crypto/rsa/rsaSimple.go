package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

func main() {
	createPrivateKey()
}

func createPrivateKey() {
	reader := rand.Reader
	bitSize := 512
	//生成密钥
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		fmt.Println("生成密钥错误", err.Error())
	}

	publicKey := key.PublicKey
	sha := sha1.New()
	msg := []byte("hello RSA")

	//公钥加密数据
	out, err := rsa.EncryptOAEP(sha, reader, &publicKey, msg, nil)
	if err != nil {
		fmt.Println("加密数据出错", err.Error())
	}

	strOut := string(out)
	fmt.Println("未加密的数据：hello RSA")
	fmt.Println("加密后的数据：", strOut)

	//保存私钥到文件
	fmt.Println("私钥：", key)
	fmt.Println("公钥：", publicKey)

	//解密消息
	out, err = rsa.DecryptOAEP(sha, reader, key, out, nil)
	if err != nil {
		fmt.Println("解密消息错误：", err.Error())
	}
	fmt.Println("解密后的消息：", string(out))
}
