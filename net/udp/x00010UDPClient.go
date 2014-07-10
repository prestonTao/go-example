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
	l, err := net.ListenPacket("udp", "127.0.0.1:0")
	checkError(err)
	defer l.Close()

	c, err := net.Dial("udp", l.LocalAddr().String())
	checkError(err)
	defer c.Close()

	ra, err := net.ResolveUDPAddr("udp", l.LocalAddr().String())
	checkError(err)
	_, err = c.(*net.UDPConn).WriteToUDP([]byte("Connection-less mode socket"), ra)
	checkError(err)
}
