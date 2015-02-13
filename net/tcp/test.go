package main

import (
	"encoding/binary"
	"fmt"
	// "io"
	"net"
	// "time"
)

var hello = []byte{22, 3, 1, 0, 186, 1, 0, 0, 182, 3, 3, 154, 202, 85, 242, 246, 106, 95, 84, 50, 80, 132, 129, 56, 78, 20,
	0, 134, 103, 243, 198, 104, 6, 109, 210, 205, 100, 216, 229, 245, 171, 253, 206, 198, 0, 0, 40, 204, 2,
	0, 204, 19, 192, 43, 192, 47, 0, 158, 192, 10, 192, 9, 192, 19, 192, 20, 192, 7, 192, 17, 0, 51, 0, 50, 0,
	57, 0, 156, 0, 47, 0, 53, 0, 10, 0, 5, 0, 4, 1, 0, 0, 101, 255, 1, 0, 1, 0, 0, 10, 0, 8, 0, 6, 0, 23, 0, 24, 0,
	25, 0, 11, 0, 2, 1, 0, 0, 35, 0, 0, 51, 116, 0, 0, 0, 16, 0, 27, 0, 25, 6, 115, 112, 100, 121, 47, 51, 8, 115,
	112, 100, 121, 47, 51, 46, 49, 8, 104, 116, 116, 112, 47, 49, 46, 49, 117, 80, 0, 0, 0, 5, 0, 5, 1, 0, 0,
	0, 0, 0, 18, 0, 0, 0, 13, 0, 18, 0, 16, 4, 1, 5, 1, 2, 1, 4, 3, 5, 3, 2, 3, 4, 2, 2, 2}

func main() {
	// Server()
	Client()

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

func Client() {
	conn, _ := net.Dial("tcp", "mandela.io:8081")
	conn.Write(hello)
}
