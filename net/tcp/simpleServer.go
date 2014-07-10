package main

import (
	"net"
	// "strconv"
	"time"
)

func main() {
	Server()
}

func Server() {
	ln, _ := net.Listen("tcp", ":8083")
	for {
		conn, _ := ln.Accept()
		time.Sleep(time.Second * 1)
		conn.Write([]byte("haha"))
		conn.Write([]byte("message from server"))
		conn.Write([]byte("message from server"))
	}

}
