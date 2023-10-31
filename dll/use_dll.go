package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	example3()
}

func example1() {
	handler, err := syscall.LoadLibrary("exportgo.dll")
	fmt.Println("DLL文件加载完成", handler, err)
	fn, err := syscall.GetProcAddress(handler, "PrintBye")
	fmt.Println("获得方法完成", fn, err)
	r1, r2, err := syscall.Syscall(fn, 0, 0, 0, 0)
	fmt.Println("执行结果", r1, r2, err)
}

func example2() {
	h := syscall.MustLoadDLL("exportgo.dll")
	c := h.MustFindProc("Sum")
	r1, r2, err := c.Call(uintptr(1), uintptr(1))
	fmt.Println(r1, r2, err)
}

func custem() {
	dll, err := syscall.LoadDLL("exportgo.dll")
	fmt.Println(dll, err)

	proc, err := dll.FindProc("PrintBye")
	fmt.Println(proc, err)
}

func example3() {
	dll, err := syscall.LoadDLL("exportgo.dll")
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println(dll, err)

	proc, err := dll.FindProc("StartUP")
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println(proc, err)

	params := "123456789"

	_, _, err = proc.Call(StrPtr(params))
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

}

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	// return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}
