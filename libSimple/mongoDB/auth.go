package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func init() {

}

func main() {
	session, err := mgo.Dial("192.168.0.40:27017")
	if err != nil {
		panic(err)
	}
	admindb := session.DB("admin")
	err = admindb.Login("root", "123456")
	if err != nil {
		fmt.Println("登录错误：", err.Error())
	}
	fmt.Println("mongodb登录成功")
}
