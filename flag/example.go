package main

import (
	"flag"
	"fmt"
)

//go run example.go init db
func main() {
	// ok := flag.String("init", "db", "create table in db")
	// flag.BoolVar(, name, value, usage)
	flag.Parse()

	fmt.Println(flag.Args())
}
