package main

import (
	"fmt"
	"os"
)

func main() {
	example1()
}

func example1() {
	file, err := os.Open("file/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(fileInfo.Size())
}
