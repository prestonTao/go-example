package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	remotAddr, err := net.ResolveUDPAddr("udp", "239.0.0.1:9982")
	if err != nil {
		log.Println("组播：组播地址格式不正确")
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   []byte{192, 168, 1, 128},
		Port: 9981,
	})
	defer conn.Close()
	if err != nil {
		log.Println("组播：监听udp出错")
	}
	_, err = conn.WriteToUDP([]byte("hhahahahah"), remotAddr)
	if err != nil {
		log.Println("组播：发送msg到组播地址出错")
	}
	read := bufio.NewReader(os.Stdin)
	for {
		b, _, err := read.ReadLine()
		if err != nil {
			log.Panic(err)
		}
		_, err = conn.WriteToUDP(b, remotAddr)
		if err != nil {
			log.Panic(err)
		}
	}
}
