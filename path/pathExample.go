package main

import (
	"fmt"
	"os"
	//"strings"
	//"unicode/utf8"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	name := example("./")
	fmt.Println(name)
}

func example1() {
	fmt.Println(filepath.FromSlash("/tao/action"))
	fmt.Println(path.Clean("/tao/action"))
	fmt.Println(path.Clean("/tao\\action"))
	fmt.Println(strings.Contains(" /tao", " ")) //任意地方包含空格都返回true
}

func example(fileName string) (name string) {
	// 遍历目录
	filepath.Walk(fileName,
		func(path string, f os.FileInfo, err error) error {

			fmt.Println(path)
			name = path

			return nil
		})
	return
}
