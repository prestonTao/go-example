package main

import (
	"bytes"
	"encoding/binary"
	"net"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// remotAddr, err := net.ResolveUDPAddr("udp", "100.64.195.116:9981")
	remotAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9981")
	chk(err)
	locaAddr, err := net.ResolveUDPAddr("udp", ":57075")
	chk(err)
	conn, err := net.ListenUDP("udp", locaAddr)
	chk(err)

	portBuf := bytes.NewBuffer([]byte{})
	binary.Write(portBuf, binary.LittleEndian, int16(9980))
	port := portBuf.Bytes()
	_, err = conn.WriteToUDP([]byte{5, 0, 0, 1, 127, 0, 0, 1, port[0], port[1]}, remotAddr)
	chk(err)

	// conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	// chk(err)
	// _, err = conn.WriteToUDP([]byte("hello world"), &net.UDPAddr{IP: net.IP{127, 0, 0, 1}, Port: 9981})
	// chk(err)
}
