package main

import (
	"fmt"
	ftp "github.com/jramnani/go-ftp"
)

func main() {
	conn()
}

func conn() {
	conn, err := ftp.Dial("localhost:21")
	if err != nil {
		t.Fatal(err)
	}
	err = conn.Login("anonymous", "anonymous@")
	if err != nil {
		t.Error(err)
	}
}
