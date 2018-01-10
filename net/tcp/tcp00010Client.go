package main

import (
	"fmt"
	"net"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	chk(err)
	_, err = conn.Write([]byte("hello"))
	chk(err)
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	chk(err)
	fmt.Println(string(buf[:n]))

}
