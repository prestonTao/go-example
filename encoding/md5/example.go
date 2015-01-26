package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	stringMD5()
	fileMD5()
}

/*
	生成一个文件的md5码
*/
func fileMD5() {
	file, err := os.Open("../../test_file/private.pem")
	if err != nil {
		panic("打开文件出错：" + err.Error())
	}
	h := md5.New()
	io.Copy(h, file)
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}

/*
	生成字符串的MD5码
*/
func stringMD5() {
	h := md5.New()
	h.Write([]byte("123456"))                          // 需要加密的字符串为 123456
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果

}
