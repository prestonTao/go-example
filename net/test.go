package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "127.0.0.1:3306")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("ok")

}
