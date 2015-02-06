package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	// "github.com/smallfish/ftp"
	// "os"
)

func main() {
	conn()
}

func conn() {
	c, err := ftp.Connect("192.186.7.152:21")
	if err != nil {
		fmt.Println("1----", err)
	}

	err = c.Login("ftpuser", "123456")
	if err != nil {
		fmt.Println("2----", err)
	}
	_, err = c.List(".")
	if err != nil {
		fmt.Println("3----", err)
	}
	c.Quit()
}

// func conn() {
// 	ftp := new(ftp.FTP)
// 	// debug default false
// 	ftp.Debug = true
// 	ftp.Connect("192.186.7.152", 21)
// 	fmt.Println(ftp.Code)

// 	// login
// 	ftp.Login("ftpuser", "123456")
// 	if ftp.Code == 530 {
// 		fmt.Println("error: login failure")
// 		os.Exit(-1)
// 	}
// 	fmt.Println(ftp.Code, "login success")

// 	// pwd
// 	ftp.Pwd()
// 	fmt.Println("code:", ftp.Code, ", message:", ftp.Message)

// 	// // make dir
// 	// ftp.Mkd("/path")
// 	// ftp.Request("TYPE I")

// 	// // stor file
// 	// b, _ := ioutil.ReadFile("/path/a.txt")
// 	// ftp.Stor("/path/a.txt", b)

// 	ftp.Quit()
// }
