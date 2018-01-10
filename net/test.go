package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:19981")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
	// _, err = conn.Write([]byte("nihao"))
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("ok")
	// }

}
