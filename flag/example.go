package main

import (
	"flag"
	"fmt"
)

//启动程序的时候，带参数
//go run example.go init db
func main() {
	// ok := flag.String("init", "db", "create table in db")
	// flag.BoolVar(, name, value, usage)
	flag.Parse()

	fmt.Println(flag.Args())
	t()
}

func t() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 || args[0] != "init" {
		fmt.Println("没有运行")
		return
	}
	fmt.Println("运行了")
}
