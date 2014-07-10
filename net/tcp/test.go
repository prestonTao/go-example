package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	// "time"
)

func main() {
	Server()
}

func Server() {
	ln, _ := net.Listen("tcp", ":8083")
	for {
		conn, _ := ln.Accept()
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(buf[:n])
		crc1 := binary.BigEndian.Uint32(buf[:4])
		crc2 := binary.LittleEndian.Uint32(buf[:4])
		fmt.Println(crc1, crc2)

		n, _ = io.ReadFull(conn, buf)
		fmt.Println(buf[:n])
		crc1 = binary.BigEndian.Uint32(buf[:4])
		fmt.Println(crc1)
	}

}
