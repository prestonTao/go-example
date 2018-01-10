package main

import (
	"net"
)

func main() {
	buf := "hello"
	c, _ := net.Dial("udp", "127.0.0.1:9981")
	c.Write([]byte(buf))

}
