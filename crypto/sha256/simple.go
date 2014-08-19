package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	simple1("a")
	simple1("taopopoo")
	simple1("tao")
	simple1("taohon")
	simple1("taohong")
	simple1("taohongfei")
}

func simple1(name string) {
	sha := sha256.New()
	sha.Write([]byte(name))
	md := sha.Sum(nil)
	fmt.Println(md)
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)

	id := new(big.Int).SetBytes(md)
	fmt.Println(id.String(), id.BitLen())
}
