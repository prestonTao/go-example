package main

import (
	"bytes"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	"golang.org/x/crypto/curve25519"
)

func main() {
	//生成A的密钥对
	prik, pubk := GenKeyPair()
	prikStr := hex.EncodeToString(prik[:])
	fmt.Println("prkA", len(prik), prikStr)

	pubkStr := hex.EncodeToString(pubk[:])
	fmt.Println("pukA", len(pubk), pubkStr)

	//生成B的密钥对
	prkB, pukB := GenKeyPair()
	fmt.Println("prkB", hex.EncodeToString(prkB[:]))
	fmt.Println("pukB", hex.EncodeToString(pukB[:]))

	var dhOutA [32]byte
	curve25519.ScalarMult(&dhOutA, &prik, &pukB)
	fmt.Println("协商密钥A", hex.EncodeToString(dhOutA[:]))

	var dhOutB [32]byte
	curve25519.ScalarMult(&dhOutB, &prkB, &pubk)
	fmt.Println("协商密钥B", hex.EncodeToString(dhOutB[:]))

}

/*
	生成公私钥对
*/
func GenKeyPair() ([32]byte, [32]byte) {
	//生成私钥
	prik := genKey()
	//生成公钥
	var pubk [32]byte
	curve25519.ScalarBaseMult(&pubk, &prik)
	return prik, pubk
}

/*
	获取一个32字节随机数作为私钥
*/
func genKey() [32]byte {
	r := randn()

	hash := sha256.New()
	hash.Write(r)
	md := hash.Sum(nil)
	var result [32]byte
	// copy(result, md)
	reader := bytes.NewReader(md)
	reader.Read(result[:])

	return result
}

//从范围[0-65535)随机选择一个数字，包括0，但不包括65535
func randn() []byte {
	max := BytesToUint64([]byte{255 - 128, 255, 255, 255, 255, 255, 255, 255})
	r := GetRandNum(int64(max))

	// timens := int64(time.Now().Nanosecond())
	// rand.Seed(timens)
	// r := rand.Uint64() //Intn(65535)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(r))
	return bs
}

/*
	获得一个随机数(0 - n]，包含0，不包含n
*/
func GetRandNum(n int64) int64 {
	if n == 0 {
		return 0
	}
	result, _ := crand.Int(crand.Reader, big.NewInt(int64(n)))
	return result.Int64()
}

//byte转uint64
func BytesToUint64(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp uint64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}
