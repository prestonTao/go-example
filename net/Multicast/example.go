package main

import (
	"net"
	"time"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	go example_send()
}

func example_send() {
	remotAddr, err := net.ResolveUDPAddr("udp", "239.255.255.250:1900")
	chk(err)

	locaAddr, err := net.ResolveUDPAddr("udp", "192.168.189.27:9980")
	chk(err)
	conn, err := net.ListenUDP("udp", locaAddr)
	chk(err)

	for {
		_, err = conn.WriteToUDP([]byte("nihao"), remotAddr)
		chk(err)
		time.Sleep(time * 5)
	}

}
