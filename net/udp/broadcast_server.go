// hello project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	udpaddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:1314")
	if err != nil {
		log.Panic(err)
	}
	addr, err := net.ResolveUDPAddr("udp", "192.168.1.128:1314")
	if err != nil {
		log.Panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	read := bufio.NewReader(os.Stdin)
	fmt.Print("自己起个名字吧：")
	b, _, err := read.ReadLine()
	if err != nil {
		log.Panic(err)
	}
	name := string(b)
	for {
		b, _, err := read.ReadLine()
		if err != nil {
			log.Panic(err)
		}
		t := time.Now()
		msg := fmt.Sprintf("[%s]：\t%s\t(%d.%d)", name, b, t.Minute(), t.Second())
		_, err = conn.WriteToUDP([]byte(msg), udpaddr)
		if err != nil {
			log.Panic(err)
		}

	}
}
