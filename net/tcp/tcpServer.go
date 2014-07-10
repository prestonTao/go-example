package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleConnection(conn net.Conn, i int) {
	fmt.Println("connect succeed! ID: ", i)
	i += 1
	time.Sleep(1 * time.Millisecond)
	_, _ = conn.Write([]byte("message from server\r\n\r\n"))
	time.Sleep(1 * time.Millisecond)
	conn.Close()
}
func main() {
	i := 0
	ln, err := net.Listen("tcp", ":8080")
	checkError(err)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			continue
		}
		i += 1
		go handleConnection(conn, i)
	}
}
