package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// c, cancel := context.WithCancel(context.Background())
	example1(nil)

	time.Sleep(time.Second * 10)
	// cancel()
	select {}
}

func example1(c context.Context) {
	if c == nil {
		fmt.Println("context is nil")
	}

	// newc1, cancel := context.WithCancel(context.Background())
	// go run1(c)
	// time.Sleep(time.Second * 5)
	// cancel()
}

// func run1(c context.Context) {
// 	go run1_child(c)
// 	count := 0
// 	for {
// 		count++
// 		select {
// 		case <-c.Done():
// 			fmt.Println("run1 done")
// 			return
// 		default:
// 		}
// 		time.Sleep(time.Second)
// 		fmt.Println("run1 haha", count)
// 	}
// }
