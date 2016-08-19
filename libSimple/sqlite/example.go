package main

import (
	"fmt"

	"github.com/sbinet/sqlite3"
)

func main() {
	db, err := sqlite3.Open("test.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, table := range db.Tables() {
		fmt.Printf(">>> table=%#v\n", table)
	}
}
