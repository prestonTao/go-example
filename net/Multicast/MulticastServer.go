package main

import (
	// "fmt"
	// "io"
	"log"
	"net"
	// "os"
)

//客户端和服务器需要在两台机器上才能实验
func main() {
	addr := &net.UDPAddr{
		IP:   net.ParseIP("224.0.1.251"),
		Port: 5353,
	}
	conn, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, 1024)
		_, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		log.Println(string(buf))
	}

}
