package main

import (
	"fmt"
	"time"
)

func main() {
	//eTest1()
	//eTest2()
	eTest3()
	//睡眠3秒钟
	time.Sleep(time.Second * 3)
}
func eTest1() {
	fmt.Println(time.Now().String())
	var customMap map[string]string = map[string]string{"tao": "nihao"}
	fmt.Println(customMap)
	fmt.Println([]string{"tao"})
}
func eTest2() {
	fmt.Println("the 1")
	tc := time.After(time.Second)
	fmt.Println("the 2")
	<-tc
	fmt.Println("the 3")
	fmt.Println(tc)
}

//计算程序运行的时间
func eTest3() {
	fmt.Println("the 1")
	t1 := time.Now()
	fmt.Println("the 2")
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Nanoseconds()) //计算出微毫秒，十亿份之一秒
}
