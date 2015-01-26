package main

import (
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2"
	"io"
	"os"
)

var MDB *mgo.Database

/*
	连接mongodb并登录
*/
func init() {
	session, err := mgo.Dial("125.64.93.83:27017")
	if err != nil {
		panic("mgo init errors")
	}
	session.SetMode(mgo.Monotonic, true)
	admindb := session.DB("admin")
	err = admindb.Login("root", "root")
	if err != nil {
		fmt.Println("登录错误：", err.Error())
	}
	fmt.Println("mongodb login success")
	MDB = session.DB("fileServer")
}

func main() {
	download()
}

/*
	上传文件
*/
func upload() {
	file, _ := os.Open("new.apk")
	gridFile := &mgo.GridFile{}
	if gridFile, err = MDB.GridFS("fs").Create("newfilename"); err != nil {
	}
	if _, err = io.Copy(gridFile, file); err != nil {
	}
	gridFile.Close()
	file.Close()
}

/*
	文件下载
*/
func download() {
	gridFile, err := MDB.GridFS("fs").Open("4fb132ccb31d490a7d00cbcf16850e")
	if err != nil {
		fmt.Println("mgo打开文件错误")
		if err == mgo.ErrNotFound {
			fmt.Println("mgo未找到错误")
		}
	}
	buf := bytes.NewBuffer([]byte{})
	io.Copy(buf, gridFile)
	file, _ := os.Create("new.apk")
	file.Write(buf.Bytes())
	file.Close()
	gridFile.Close()
}
