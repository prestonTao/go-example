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
	// fileMD5()
	// parseMD5()
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
	h.Write([]byte("123456")) // 需要加密的字符串为 123456
	fmt.Println(h.Sum(nil))
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果

	bs, _ := hex.DecodeString(hex.EncodeToString(h.Sum(nil)))
	fmt.Println(bs)

}

//d751713988987e9331980363e24189ce
func parseMD5() {
	strMD5 := "d751713988987e9331980363e24189ce"
	fmt.Println([]byte(strMD5))

}
