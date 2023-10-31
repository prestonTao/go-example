package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c, cancel := context.WithCancel(context.Background())
	go example1(c)

	time.Sleep(time.Second * 10)
	cancel()
	select {}
}

func example1(c context.Context) {
	newc1, cancel := context.WithCancel(context.Background())
	go run1(c)
	go run2(newc1)
	time.Sleep(time.Second * 5)
	cancel()
}

func run1(c context.Context) {
	go run1_child(c)
	count := 0
	for {
		count++
		select {
		case <-c.Done():
			fmt.Println("run1 done")
			return
		default:
		}
		time.Sleep(time.Second)
		fmt.Println("run1 haha", count)
	}
}

func run1_child(c context.Context) {
	count := 0
	for {
		count++
		select {
		case <-c.Done():
			fmt.Println("run1_child done")
			return
		default:
		}
		time.Sleep(time.Second)
		fmt.Println("run1_child haha", count)
	}
}

func run2(c context.Context) {
	count := 0
	for {
		count++
		select {
		case <-c.Done():
			fmt.Println("run2 done")
			return
		default:
		}
		time.Sleep(time.Second)
		fmt.Println("run2 haha", count)
	}
}
