package main

import (
	"fmt"
	"time"
)

var container chan bool

func main() {
	var arrays []int = []int{25, 12, 36, 24, 96, 54, 28}
	container = make(chan bool, len(arrays))
	for i := 0; i < len(arrays); i++ {
		go tosleep(arrays[i])
	}
	for i := 0; i < len(arrays); i++ {
		<-container
	}
}

func tosleep(data int) {
	time.Sleep(time.Duration(data))
	fmt.Println(data)
	container <- true
}
