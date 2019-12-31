package main

import (
	"bytes"
	"encoding/hex"

	"golang.org/x/crypto/ed25519"

	// "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
)

func main() {
	hash := sha256.New

	master := []byte{0x00, 0x01, 0x02, 0x03} // i.e. NOT this.

	// salt := make([]byte, hash().Size())
	// n, err := io.ReadFull(rand.Reader, salt)
	// if n != len(salt) || err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	salt := sha256.Sum256(master)

	// info := []byte{0x03, 0x14, 0x15, 0x92, 0x65}

	// Create the key derivation function
	hkdf := hkdf.New(hash, master, salt[:], nil)

	// Generate the required keys
	keys := make([][]byte, 5)
	for i := 0; i < len(keys); i++ {
		keys[i] = make([]byte, 32)
		n, err := io.ReadFull(hkdf, keys[i])
		if n != len(keys[i]) || err != nil {
			fmt.Println("error:", err)
			return
		}
	}

	fmt.Println(keys)

	temp := bytes.NewBuffer(keys[0])
	puk, prk, _ := ed25519.GenerateKey(temp)
	fmt.Println("打印密钥\n", hex.EncodeToString(puk), "\n", hex.EncodeToString(prk))

}
