package main

import (
	"../../scrypt"
	"fmt"
	"math/big"
)

func main() {
	dk, err := scrypt.Key([]byte("some password"), []byte("tao"), 16384, 8, 1, 32)
	dkInt := new(big.Int).SetBytes(dk)
	fmt.Println(err)
	fmt.Println("长度：", dkInt.BitLen(), "内容：", dkInt.String())
}
