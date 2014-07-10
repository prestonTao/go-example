package main

import (
	"fmt"
	// "io"
	"net"
	"os"
	// "time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkError(err)
 	buf := make([]byte,1024) 
	conn.Write([]byte("I'm client"))

	
	n, err := conn.Read(buf)
	checkError(err)
	fmt.Println(string([]byte(buf[:n])))

}
