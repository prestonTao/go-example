package main

import (
	"fmt"
	"github.com/dchest/scrypt"
)

func main() {
	dk, err := scrypt.Key([]byte("some password"), []byte("tao"), 16384, 8, 1, 32)
	fmt.Println(err)
	fmt.Println(string(dk))
}
