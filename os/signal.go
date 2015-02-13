package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	example1()
}

func example2() {
	fmt.Println("listen")
	c := make(chan os.Signal)
	signal.Notify(c)
	sig := <-c

	fmt.Println("------", sig)
}

/*
	选择监听程序退出型号量
*/
func example1() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)
}
