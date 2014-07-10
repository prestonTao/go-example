package main

// //添加一些编译选项

// #cgo CFLAGS:  -DPNG_DEBUG=1
// #cgo linux CFLAGS:  -DLINUX=1
// #cgo LDFLAGS:  -lpng
// #include <png.h>
import "C"

func Hello() {
	C.hello()
}

func main() {
	Hello()
}
