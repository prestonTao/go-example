package main

import (
	fm "fmt"
	"net/url"
	rt "runtime"
)

func main() {
	example1()
}

func example() {
	var urlStr string = "http://baidu.com/index.php/?abc=1_羽毛"
	l, err := url.ParseQuery(urlStr)
	fm.Println(l, err)
	l2, err2 := url.ParseRequestURI(urlStr)
	fm.Println(l2, err2)

	l3, err3 := url.Parse(urlStr)
	fm.Println(l3, err3)
	fm.Println(l3.Path)
	fm.Println(l3.RawQuery)
	fm.Println(l3.Query())
	fm.Println(l3.Query().Encode())

	fm.Println(l3.RequestURI())
	fm.Printf("Hello World! version : %s", rt.Version())
}

func example1() {
	urlStr := "http://baidu.com/index.php/?abc=1_羽毛"
	l, _ := url.ParseQuery(urlStr)
	fm.Println(l.Encode())
}
