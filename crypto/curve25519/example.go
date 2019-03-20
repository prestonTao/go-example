package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/curve25519"
)

func main() {

	prik, pubk := GenKeyPair()
	prikStr := hex.EncodeToString(prik[:])
	fmt.Println("prikey", len(prik), prikStr)

	pubkStr := hex.EncodeToString(pubk[:])
	fmt.Println("pubkey", len(pubk), pubkStr)

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
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	r := rand.Uint64() //Intn(65535)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, r)
	return bs
}
