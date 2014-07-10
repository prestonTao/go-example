package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"sync"
)

// key
const aesTable = "abcdefghijklmnopkrstuvwsyz012345"

var (
	block cipher.Block
	mutex sync.Mutex
)

func main() {
	// encript
	encodeBytes, err := Encrypt([]byte("weloveyouchinese"))
	if err != nil {
		fmt.Println("Encrypt: ", err)
		return
	}

	// encode println
	fmt.Printf("Encrypt code: %x\n", string(encodeBytes))

	// decrypt
	decodeBytes, err := Decrypt(encodeBytes)
	if err != nil {
		fmt.Println("Decrypt: ", err)
		return
	}

	// decode println
	fmt.Println("Decrypt code: ", string(decodeBytes))
}

// AES加密
func Encrypt(src []byte) ([]byte, error) {
	// 验证输入参数
	// 必须为aes.Blocksize的倍数
	if len(src)%aes.BlockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	encryptText := make([]byte, aes.BlockSize+len(src))

	iv := encryptText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(encryptText[aes.BlockSize:], src)

	return encryptText, nil
}

// AES解密
func Decrypt(src []byte) ([]byte, error) {
	// hex
	decryptText, err := hex.DecodeString(fmt.Sprintf("%x", string(src)))
	if err != nil {
		return nil, err
	}

	// 长度不能小于aes.Blocksize
	if len(decryptText) < aes.BlockSize {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	iv := decryptText[:aes.BlockSize]
	decryptText = decryptText[aes.BlockSize:]

	// 验证输入参数
	// 必须为aes.Blocksize的倍数
	if len(decryptText)%aes.BlockSize != 0 {
		return nil, errors.New("crypto/cipher: ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(decryptText, decryptText)

	return decryptText, nil
}

func init() {
	mutex.Lock()
	defer mutex.Unlock()

	if block != nil {
		return
	}

	cblock, err := aes.NewCipher([]byte(aesTable))
	if err != nil {
		panic("aes.NewCipher: " + err.Error())
	}

	block = cblock
}
