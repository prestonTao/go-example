package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// readTextFile("class_old.txt")
	createDIR("class_new.txt")
}

func createDIR(path string) {
	file, _ := os.Open(path)
	buf := bufio.NewReader(file)
	count := 0
	for {
		count++
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		if string(line) == "" {
			continue
		}
		names := strings.Split(string(line), "||")
		if len(names) < 6 {
			panic(count)
		}
		familia := strings.Join(names[:2], "_")
		genus := strings.Join(names[2:4], "_")
		temps := strings.Split(names[4], "/")
		plant := strings.Join(temps, "&") + "_" + names[5]
		os.MkdirAll(filepath.Join("root", familia, genus, plant), 0777)

		// name := strings.Join(names[:5], "||")
		// name = name + "||" + strings.Join(names[5:], " ")
		// fmt.Println(names)
		// newfile.WriteString(name + "\n")
	}
	file.Close()
}

//一行一行的读文本文件，并打印出来
func readTextFile(path string) {
	newfile, _ := os.Create("class_new.txt")

	file, _ := os.Open(path)
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		if string(line) == "" {
			continue
		}
		names := strings.Split(string(line), "||")
		if len(names) < 6 {
			panic(names)
		}
		name := strings.Join(names[:5], "||")
		name = name + "||" + strings.Join(names[5:], " ")
		fmt.Println(names)
		newfile.WriteString(name + "\n")
	}
	file.Close()
}
