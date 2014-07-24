package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	example2()
}

func example1() {
	path := "file/data_1"

	err := os.Truncate(path, int64(100))
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err.Error())
	}

	size, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(size)

	buf := bytes.NewBuffer([]byte("t"))
	fmt.Println(buf.Bytes())
	file.WriteAt(buf.Bytes(), 40)

}

func example2() {
	buf := bytes.NewBuffer([]byte("taopopootaohongefifnlvnknihaoa"))
	fmt.Println(buf.Bytes(), buf.Len())
	result, e := binary.Varint(buf.Bytes())
	fmt.Println(result, e)
}
