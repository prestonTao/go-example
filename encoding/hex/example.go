package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

//将uint16值转化为16进制字符串
func main() {
	var h uint16 = 57553
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, h)
	str := hex.EncodeToString(buf.Bytes())
	fmt.Println(str)
}
