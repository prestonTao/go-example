package main

import (
	"log"
	"net"
)

func main() {
	startUP()
}

type Server struct {
	conn net.Conn
}

func (this *Server) read() {
	buf := make([]byte, 1024)
	n, _ := this.conn.Read(buf)
	log.Println(string(buf[:n]))
}
func (this *Server) write() {

}

func startUP() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		s := Server{conn: conn}
		go s.read()
	}
}
