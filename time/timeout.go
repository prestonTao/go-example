package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go TestChan(c)
	str := <-c
	fmt.Println(str)

	time.Sleep(time.Second * 5)
}

func TestChan(c chan string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("超时了")
		} else {
			fmt.Println("没超时")
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("没超时")
			} else {
				fmt.Println("超时了")
			}
		}()
		time.Sleep(time.Second * 2)
		c <- "ha"
		close(c)
	}()
	time.Sleep(time.Second * 3)
	c <- "nihao"
	close(c)
}
