package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

var sha, sha2 hash.Hash

func init() {
	sha = sha256.New()
	sha2 = sha256.New() // hash twice
}

func Hash(ba []byte) []byte {
	sha.Reset()
	sha2.Reset()
	ba = sha.Sum(ba)
	return sha2.Sum(ba)
}
func main() {
	bytes := Hash([]byte(`prestonprestonprestonprestonprestonprestonpr
		estonprestonprestonprestonprestonprestonp
		restonprestonprestonprestonprestonprestonpreston`))
	fmt.Println(string(bytes))

	//---------------------------------------
	hash := sha256.New()
	hash.Write([]byte("a"))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)
}
