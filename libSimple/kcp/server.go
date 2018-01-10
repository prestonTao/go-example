package main

import (
	"fmt"
	"log"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	start()
}

func start() {

	fmt.Println("start server")
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}
	log.Println("listening on:", lis.Addr())
	for {
		buf := make([]byte, 1024)
		ss, err := lis.AcceptKCP()
		if err != nil {
			panic(err)
		}
		fmt.Println("new conn")
		n, err := ss.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
