package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// l, err := net.ListenPacket("udp", "127.0.0.1:9981")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer l.Close()

	// c, err := net.Dial("udp", l.LocalAddr().String())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer c.Close()

	// ra, err := net.ResolveUDPAddr("udp", l.LocalAddr().String())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = c.(*net.UDPConn).WriteToUDP([]byte("Connection-oriented mode socket"), ra)
	// if err == nil {
	// 	log.Fatal(err)
	// }

	//-----------------------------------------------------------------------------------
	//-----------------------------------------------------------------------------------

	// addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9981")
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
	count, remoteAddr, err := conn.ReadFromUDP(buf)
	if err != nil || count < 1 {
		log.Fatal("ReadFromUDP error ", err)
	}
	fmt.Println(string(buf), "\n  |  ", remoteAddr.IP, "  |  ", remoteAddr.Port)

	//-----------------------------------------------------------------------------------
	//-----------------------------------------------------------------------------------

	//server := NewUDPServer("0:9981")
	// buf := make([]byte, 1024)
	// addr, err := net.ResolveUDPAddr("udp", ":9981")
	// sock, err := net.ListenUDP("udp", addr)
	// count, remoteAddr, err := sock.ReadFromUDP(buf)
	// if err != nil || count < 1 {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(buf), "  |  ", remoteAddr.IP, "  |  ", remoteAddr.Port)

}
