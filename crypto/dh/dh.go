/*
	DH密钥交换和ECDH原理
	http://www.tuicool.com/articles/em6zEb
	非对称加密算法-DH算法
	http://blog.csdn.net/kongqz/article/details/6302913
*/

package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	example()
}

func example() {
	rand.Seed(int64(time.Now().Nanosecond()))

	p := rand.Intn(100)
	g := rand.Intn(100)
	fmt.Println("P: ", p, "  G: ", g)

	alice_rand := rand.Intn(100)
	bob_rand := rand.Intn(100)
	fmt.Println("Alice rand: ", alice_rand, "  Bob rand: ", bob_rand)

	alice_a := p
	for i := 0; i < alice_rand; i++ {
		alice_a = alice_a + ((alice_a * g) & 0xffffffffffffffff)
	}
	alice_a = alice_a % p
	fmt.Println("Alice g^a mod p: ", alice_a, g^alice_rand)
}
