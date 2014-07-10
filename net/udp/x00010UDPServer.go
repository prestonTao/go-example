package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	sock, err := net.ListenUDP("udp", addr)
	checkError(err)
	defer sock.Close()
	// for{

	// }
	fmt.Println(sock.IP, "  |  ", sock.Port)
	buf := make([]byte, 1024)
	_, remote, err := sock.ReadFromUDP(buf)
	checkError(err)
	fmt.Println(remote.IP, "  |  ", remote.Port)
}
