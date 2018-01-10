package main

import (
	// "log"
	"net"
	// "strconv"
	"time"
)

func main() {
	Client()
}

func Client() {
	conn, _ := net.Dial("tcp", "192.168.1.210:9981")
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
