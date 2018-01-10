package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(r.Intn(100))
	// }

	// x := rand.NewSource(9)

	// fmt.Println(rand.NewSource(9))
	example2()
}

//从范围[0-65535)随机选择一个数字，包括0，但不包括65535
func randn() {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	r := rand.Intn(65535)
	fmt.Println(r)

}

func example() {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}

/*
	产生[0-5)的数字随机排序
*/
func example2() {
	for i := 0; i < 100; i++ {
		result := rand.Perm(5)
		fmt.Println(result)
	}
}
