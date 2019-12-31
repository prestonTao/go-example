package main

import (
	"fmt"
	// "log"
	"net"
	// "strconv"
	"time"
)

func main() {
	Client()
}

func Client() {
	conn, err := net.Dial("tcp", "120.79.179.22:19981")
	fmt.Println(err)
	// buf := make([]byte, 1024)
	// n, _ := conn.Read(buf)
	// log.Println(buf[:n])
	// log.Println(string(buf[:n]))
	// buf = make([]byte, 1024)
	// conn.Read(buf)
	// log.Println(string(buf))

	conn.Write([]byte("message from server\r\n\r\n"))
	time.Sleep(time.Second * 3)

}
