package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9981")
	addr, err := net.ResolveTCPAddr("tcp", ":9981")
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	if err != nil || count < 1 {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

}
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9981")
	addr, err := net.ResolveTCPAddr("tcp", ":9981")
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	if err != nil || count < 1 {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

}
