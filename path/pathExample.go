package main

import (
	"fmt"
	//"strings"
	//"unicode/utf8"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(filepath.FromSlash("/tao/action"))
	fmt.Println(path.Clean("/tao/action"))
	fmt.Println(path.Clean("/tao\\action"))
	fmt.Println(strings.Contains(" /tao", " ")) //任意地方包含空格都返回true
}
