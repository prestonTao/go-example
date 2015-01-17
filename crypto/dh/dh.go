/*
	DH算法

	DH算法的出现就是用来进行密钥传输的。DH算法是基于离散对数实现的。用户A和B如何利用RSA算法来传输密钥？
	在通信前，用户A和B双方约定2个大整数n和g,其中1<g<n，这两个整数可以公开
	1)   A随机产生一个大整数a，然后计算Ka=ga mod n。（a需要保密）
	2)   B随机产生一个大整数b，然后计算Kb=gb mod n。（b需要保密）
	3)   A把Ka发送给B,B把Kb发送给A
	4)   A计算K=Kba mod n
	5)   B计算K=Kab mod n
	由于Kba mod n= （gb mod n）a mod n= （ga mod n）b mod n，因此可以保证双方得到的K是相同的，K即是共享的密钥。

	DH密钥交换和ECDH原理
	http://www.tuicool.com/articles/em6zEb
	非对称加密算法-DH算法
	http://blog.csdn.net/kongqz/article/details/6302913
*/
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	// These 2 groups are used in SSH and provide good builtin groups

	// diffie-hellman-group1-sha1
	group1p = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF"

	// diffie-hellman-group14-sha1
	group14p = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF"
)
var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)
var bigTwo = big.NewInt(2)

func main() {
	example()
}

func example() {
	/*
	   1.爱丽丝与鲍伯协定使用p=23以及g=5.
	*/
	p, _ := new(big.Int).SetString(group14p, 16)
	g := new(big.Int).SetInt64(2)
	fmt.Println("p: ", p)
	fmt.Println("g: ", g)

	/*
		2.爱丽丝选择一个秘密整数a=6, 计算A = g^a mod p并发送给鲍伯。
			A = 5^6 mod 23 = 8.
	*/
	q := new(big.Int).Sub(p, bigOne)
	q.Div(q, bigTwo)
	a, err := rand.Int(rand.Reader, q)
	if err != nil {
		panic("error: " + err.Error())
	}
	A := new(big.Int).Exp(g, a, p)

	/*
		3.鲍伯选择一个秘密整数b=15, 计算B = g^b mod p并发送给爱丽丝。
			B = 5^15 mod 23 = 19.
	*/
	b, err := rand.Int(rand.Reader, q)
	if err != nil {
		panic("error: " + err.Error())
	}
	B := new(big.Int).Exp(g, b, p)

	/*
		4.爱丽丝计算s = B a mod p
		  19^6 mod 23 = 2.
	*/
	s1 := new(big.Int).Exp(B, a, p)
	/*
		5.鲍伯计算s = A b mod p
		   8^15 mod 23 = 2.
	*/
	s2 := new(big.Int).Exp(A, b, p)

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)

}
