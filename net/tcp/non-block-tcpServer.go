package main

import (
	"fmt"
	"net"
	"os"
	// "time"
)


type Server struct{
	conn net.Conn
	outCtrl chan string
}

func (this *Server) read(){
	for{
		buf := make([]byte, 1024)
		n,err := this.conn.Read(buf)
		checkError(err)
		fmt.Println("收到消息：",string([]byte(buf[:n])))
		this.outCtrl <- string([]byte(buf[:n]))
	}
}

func (this *Server) write(){
	for{
		select {
		case str := <- this.outCtrl:
			fmt.Println("准备写入")
			_,err := this.conn.Write([]byte(str))
			checkError(err)
			fmt.Println("写入完成")
		}
	}
}

func Listen(){
	ln, err := net.Listen("tcp", ":8080")
	checkError(err)
	for{
		conn, err := ln.Accept()
		checkError(err)
		s := Server{conn : conn, outCtrl : make(chan string)}
		go s.read();
		go s.write();
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


func main() {
	Listen()
}
