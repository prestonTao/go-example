package main

import (
	"bufio"
	"log"
	"os"
	"syscall"
)

//监听命令行输入，并把输入内容打印出来
func main() {
	running := true
	reader := bufio.NewReader(os.Stdin, syscall.SIGINT)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}
		log.Println("command", command)
	}
}
