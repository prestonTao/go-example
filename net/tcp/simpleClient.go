package main

import (
	"log"
	"net"
	// "strconv"
	"time"
)

func main() {
	Client()
}

func Client() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8083")
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	log.Println(buf[:n])
	log.Println(string(buf[:n]))
	time.Sleep(time.Second * 1)
	buf = make([]byte, 1024)
	conn.Read(buf)
	log.Println(string(buf))
}
