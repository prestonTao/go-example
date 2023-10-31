package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	go server()
	client()
}

func server() {
	i := 0
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			continue
		}
		i += 1
		go handleConnection(conn, i)
	}
}

func handleConnection(conn net.Conn, i int) {
	fmt.Println("connect succeed! ID: ", i)
	i += 1
	// time.Sleep(1 * time.Millisecond)
	conn.SetWriteDeadline(time.Now())

	zeroTime := time.Time{}
	fmt.Println(zeroTime.IsZero())

	conn.SetWriteDeadline(zeroTime)
	_, err := conn.Write([]byte("message from server"))
	checkError(err)
	var buf [512]byte
	// conn.SetReadDeadline(time.Now())
	_, err = conn.Read(buf[:])
	checkError(err)

	// time.Sleep(1 * time.Millisecond)
	conn.Close()
}

//-------------------------

func client() {
	for i := 0; i < 1; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		checkError(err)

		fmt.Println("connect succeed! ID:", i+1)
		var buf [512]byte
		for {
			_, err := conn.Read(buf[0:])
			if err != nil {
				if err == io.EOF {
					break
				}
				checkError(err)
			}
		}
		fmt.Println(string(buf[:]))
		// time.Sleep(1 * time.Millisecond)
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
