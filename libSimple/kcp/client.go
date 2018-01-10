package main

import (
	"fmt"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	start()
}

func start() {

	fmt.Println("start client")
	kcpconn, err := kcp.DialWithOptions("192.168.1.210:10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}
	n, err := kcpconn.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	//	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	//	for {
	//		buf := new([]byte, 1024)
	//		ss, err := lis.AcceptKCP()
	//		n, err := ss.Read(buf)
	//		fmt.Println(string(buf[:n]))
	//	}
}
