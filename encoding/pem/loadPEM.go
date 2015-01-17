package main

import (
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
)

func main() {
	example()
}

func example() {
	data, err := ioutil.ReadFile("dh512.pem")
	if err != nil {
		fmt.Println("open file error")
	}
	block, _ := pem.Decode(data)

	val := new(Group)
	x, err := asn1.Unmarshal(block.Bytes, val)
	fmt.Println(x)
	fmt.Println(block.Bytes)
	fmt.Println(block)

	fmt.Println(val)

}

type Group struct {
	P *big.Int
	G *big.Int
}
