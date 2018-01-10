package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("start")
	addr, err := net.ResolveUDPAddr("udp", ":9981")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		//send
		bs := []byte("hello")
		conn.Write(bs)
		//recv
		count, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil || count < 1 {
			log.Fatal("ReadFromUDP error ", err)
		}
		fmt.Println(string(buf[:count]), "\n  |  ", remoteAddr.IP, "  |  ", remoteAddr.Port)
	}

}
