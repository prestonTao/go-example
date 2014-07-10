package main

import (
	// "fmt"
	// "io"
	"log"
	"net"
	// "os"
)

func main() {
	remotAddr, err := net.ResolveUDPAddr("udp", "224.0.0.251:5353")
	if err != nil {
		log.Println("组播：组播地址格式不正确")
	}
	locaAddr, err := net.ResolveUDPAddr("udp", "192.168.1.103:9999")
	if err != nil {
		log.Println("组播：本地ip地址格式不正确")
	}
	conn, err := net.ListenUDP("udp", locaAddr)
	defer conn.Close()
	if err != nil {
		log.Println("组播：监听udp出错")
	}
	_, err = conn.WriteToUDP([]byte("hhahahahah"), remotAddr)
	if err != nil {
		log.Println("组播：发送msg到组播地址出错")
	}

}

// func main() {
// 	addr := &net.UDPAddr{
// 		IP:   net.ParseIP("224.0.0.251"),
// 		Port: 5353,
// 	}
// 	conn, err := net.ListenMulticastUDP("udp4", nil, addr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// for {
// 	// 	buf := make([]byte, 1024)
// 	// 	_, _, err := conn.ReadFromUDP(buf)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	log.Println(string(buf))
// 	// }

// 	_, err = conn.WriteToUDP([]byte("hahhahahahaha"), addr)
// 	if err != nil {
// 		panic(err)
// 	}

// }
