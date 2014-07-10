package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	sample1()
	// sample2()
	fmt.Println("-------------------------")
	sample3()
}

func sample1() {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("%x", buf.Bytes())
}
func sample2() {
	var a uint = 0xbeefcafe
	fmt.Println(*(*byte)(unsafe.Pointer(&a)))
}

func sample3() {
	b := []byte{0x00, 0x00, 0x03, 0xe8}
	buf := bytes.NewBuffer(b)
	var x int32
	binary.Read(buf, binary.BigEndian, &x)
	fmt.Println(x)
	fmt.Println(strings.Repeat("-", 100))
	x = 1000
	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, x)
	fmt.Println(buf.Bytes())

}
