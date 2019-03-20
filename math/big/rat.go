package main

import (
	"fmt"
	"math/big"
)

/*
   a := 1.69
   b := 1.7
   c := a * b     //结果应该是2.873
   fmt.Println(c) //输出的是2.8729999999999998
*/
func main() {

	//	big.NewFloat(1).
	one, ok := big.NewRat(169, 100).Float64()
	fmt.Println(one, ok)

	two, ok := new(big.Rat).Mul(big.NewRat(169, 100), big.NewRat(17, 10)).Float64()
	fmt.Println(two, ok)

}
