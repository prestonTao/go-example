package main

import (
	"fmt"
	"math/big"
)

/*
	这个例子计算任然不精确
*/
func main() {
	one := big.NewFloat(float64(1.234567890))
	value, acc := one.Float64()
	fmt.Println(value, acc)

	//    a := 1.69
	//    b := 1.7
	//    c := a * b     //结果应该是2.873
	//    fmt.Println(c) //输出的是2.8729999999999998
	value, acc = big.NewFloat(0).Mul(big.NewFloat(1.69), big.NewFloat(1.7)).Float64()
	fmt.Println(value, acc)

	//	big.NewFloat(0).Mul(big.NewFloat(1.69), big.NewFloat(1.7)).Format()
}
