package main

import (
	"fmt"
	//"strings"
	//"unicode/utf8"
	//"path"
	//"path/filepath"
	//"strings"
	"errors"
)

func main() {
	e := GetUser()
	fmt.Println(e)
	err := e.(UserError)
	fmt.Println(err.name)
	fmt.Println(errors.New("这是新添加的一个error"))
}

type UserError struct {
	name    string
	age     int
	content string
}

func (e UserError) Error() string {
	return e.content
}

func GetUser() (e error) {
	e = UserError{"tao", 22, "名称错误"}
	return
}
