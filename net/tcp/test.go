package main

import (
	"encoding/binary"
	"fmt"
	// "io"
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
		fmt.Println("接收总长度：", n)
		fmt.Println("接收内容:\n", buf[:n], "\n")
		fmt.Println("Type:", buf[0:1])
		// sizeBytes := append([]byte{0,buf[1],buf[2],buf[3],buf[4]}, buf[1:4]...)
		size := binary.BigEndian.Uint32([]byte{0, buf[1], buf[2], buf[3], buf[4]})
		fmt.Println("length:", size)
		// crc1 := binary.BigEndian.Uint32(buf[:4])
		// crc2 := binary.LittleEndian.Uint32(buf[:4])
		// fmt.Println(crc1, crc2)

		// n, _ = io.ReadFull(conn, buf)
		// fmt.Println(buf[:n])
		// crc1 = binary.BigEndian.Uint32(buf[:4])
		// fmt.Println(crc1)
	}

}
