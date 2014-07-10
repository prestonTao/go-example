package main

import "net"

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9981")
	chk(err)
	_, err = conn.Write([]byte("hello"))
	chk(err)

}
