// goroutine project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	//c := make(chan int, 1)
	//select{
	//	case v := <- c:
	//		fmt.Println(v)
	//	case <-
	//}
	foo()
}
func foo() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(1 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

func too() {
	c := make(chan bool)
	timeout := make(chan bool)
	go func() {
		time.Sleep(1e9)
		c <- true
	}()
	select {
	case <-c:
	case <-timeout:
	}
}
