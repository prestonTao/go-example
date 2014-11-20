package main

import (
	"fmt"
	"time"
)

func main() {
	example1()
}

func example1() {
	str := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(str)
}
