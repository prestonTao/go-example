package main

import (
	"fmt"
	// "time"
	//	"reflect"
)

func main() {
	//	c := make(chan int)
	//	putChan(c)
	// time.Sleep(time.Second * 10)
	// content := <-c
	// fmt.Println(content)
	set()
}

//func putChan(chans chan int) {
//	for chanone := range chans {
//		fmt.Println(reflect.TypeOf(chanone))

//	}
//}

func set() {
	c := make(chan int, 1)
	close(c)
	done := false
	select {
	case c <- 1:
		fmt.Println("11111111")
		done = true
	default:
		fmt.Println("22222222")
		done = false
	}
	if done {
		fmt.Println("未关闭")
	} else {
		fmt.Println("关闭了")
	}
}
