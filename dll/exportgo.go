/*
编译成 DLL 文件：
go build -buildmode=c-shared -o exportgo.dll exportgo.go
编译后得到 exportgo.dll 和 exportgo.h两个文件。
*/
package main

import (
	"fmt"
)

import "C"

// import "fmt"

//export PrintBye
func PrintBye() {
	// fmt.Println("From DLL: Bye!")
}

//export Sum
func Sum(a int, b int) int {
	// fmt.Println("hello")
	return a + b
}
func main() {
	// Need a main function to make CGO compile package as C shared library
}

//export StartUP
func StartUP(pwd *C.char) {
	fmt.Println("StartUP", C.GoString(pwd))
}
