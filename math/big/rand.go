package main

import (
	// crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	randInt := RandNodeId(512)
	fmt.Println(randInt.String())

	randOne()
}

//随机得到比特位数的整数
func RandNodeId(size int) *big.Int {
	min := rand.New(rand.NewSource(99))
	timens := int64(time.Now().Nanosecond())
	min.Seed(timens)
	maxId := new(big.Int).Lsh(big.NewInt(1), uint(size))
	randInt := new(big.Int).Rand(min, maxId)
	return randInt
}

func randOne() {
	// crand.Int(crand.Reader,)
	base, _ := new(big.Int).SetString("8", 10)
	fmt.Println(base.Exp(x, y, m))
}
