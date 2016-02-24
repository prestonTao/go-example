// hello project main.go
package main

import (
	// "bufio"
	"fmt"
	"log"
	"net"
	// "os"
	// "time"
)

func main() {
	// udpaddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:1314")
	// if err != nil {
	// 	log.Panic(err)
	// }
	addr, err := net.ResolveUDPAddr("udp", "192.168.1.128:9982")
	if err != nil {
		log.Panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	// read := bufio.NewReader(os.Stdin)
	// fmt.Print("自己起个名字吧：")
	// b, _, err := read.ReadLine()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// name := string(b)
	for {
		var b [512]byte
		n, _, err := conn.ReadFromUDP(b[:])
		if err != nil {
			log.Panic(err)
		}
		if n != 0 {
			fmt.Printf("%s\n", b[0:n])
		}
	}
}
