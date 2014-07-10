package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	f := Foo{
		Name: "tao",
		Age:  39,
	}
	buf := bytes.NewBuffer([]byte{})
	// binary.Read(buf, binary.BigEndian, f)
	binary.Write(buf, binary.BigEndian, f)
	fmt.Println("---------------")
	fmt.Println(buf)
	fmt.Println("---------------")
}

type Foo struct {
	Name string
	Age  int32
}
