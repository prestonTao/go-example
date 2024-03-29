package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var key = []byte{
	0x2b, 0x7e, 0x15, 0x16,
	0x28, 0xae, 0xd2, 0xa6,
	0xab, 0xf7, 0x15, 0x88,
	0x09, 0xcf, 0x4f, 0x3c,
}

//key = []byte{}

var src = []byte{
	0x32, 0x43, 0xF6, 0xA8, 0x88, 0x5A, 0x30, 0x8D,
	0x31, 0x31, 0x98, 0xA2, 0xE0, 0x37, 0x07, 0x34,
}

var ac = []byte{
	0x39, 0x25, 0x84, 0x1D, 0x02, 0xDC, 0x09, 0xFB,
	0xDC, 0x11, 0x85, 0x97, 0x19, 0x6A, 0x0B, 0x32,
}

func main() {
	//	a := "hello"
	//	src = []byte(a)
	//	lenhth :=
	fmt.Println("key ", key)
	fmt.Println("src ", src)

	arrEncrypt, _ := Encrypt(key, src)
	fmt.Println("ac ", ac)
	fmt.Println("加密后 ", arrEncrypt)
	strMsg, _ := Decrypt(key, arrEncrypt)
	fmt.Println("src ", strMsg)
}

//加密字符串
func Encrypt(key, srcByte []byte) ([]byte, error) {
	//	fmt.Println(aes.BlockSize)
	//	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(srcByte))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCTR(aesBlockEncrypter, key)
	aesEncrypter.XORKeyStream(encrypted, srcByte)
	//	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	//	aesEncrypter.XORKeyStream(encrypted, srcByte)
	//	aesBlockEncrypter.Encrypt(encrypted, srcByte)
	return encrypted[:], nil
}

//解密字符串
func Decrypt(key, src []byte) (strDesc []byte, err error) {

	//	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	//	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	//	aesDecrypter.XORKeyStream(decrypted, src)
	aesBlockDecrypter.Decrypt(decrypted, src)
	return decrypted, nil
}

/*
	随机获取一个128位的key
*/
func RandKey128() []byte {
	min := rand.New(rand.NewSource(99))
	min.Seed(int64(time.Now().Nanosecond()))
	maxId := new(big.Int).Lsh(big.NewInt(1), 128)
	randInt := new(big.Int).Rand(min, maxId)
	return randInt.Bytes()
}
