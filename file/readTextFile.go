package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readTextFile("index-.html")
}

//一行一行的读文本文件，并打印出来
func readTextFile(path string) {
	file, _ := os.Open(path)
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
	}
	file.Close()
}
