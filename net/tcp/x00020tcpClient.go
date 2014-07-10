package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("udp", "192.168.1.107:8080")
		checkError(err)

		fmt.Println("connect succeed! ID:", i+1)
		var buf [512]byte
		for {
			_, err := conn.Read(buf[0:])
			if err != nil {
				if err == io.EOF {
					break
				}
			}
		}
		fmt.Println(string(buf[:]))
		time.Sleep(1 * time.Millisecond)
		conn.Close()
	}
}
