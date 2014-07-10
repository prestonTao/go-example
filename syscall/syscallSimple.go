package main

import (
	"log"
	"syscall"
	"unsafe"
)

func main() {
	sample1()
	sample2()
}

func sample1() {
	//首先,准备输入参数, GetDiskFreeSpaceEx需要4个参数, 可查MSDN
	// dir := "C:"
	lpFreeBytesAvailable := int64(0) //注意类型需要跟API的类型相符
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)

	//获取方法的引用
	kernel32, err := syscall.LoadLibrary("Kernel32.dll") // 严格来说需要加上 defer syscall.FreeLibrary(kernel32)
	if err != nil {

	}
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	//执行之. 因为有4个参数,故取Syscall6才能放得下. 最后2个参数,自然就是0了
	r, _, errno := syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	log.Println(errno)
	// 注意, errno并非error接口的, 不可能是nil
	// 而且,根据MSDN的说明,返回值为0就fail, 不为0就是成功
	if r != 0 {
		log.Printf("Free %dmb", lpTotalNumberOfFreeBytes/1024/1024)
	}
}

func sample2() {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("F:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	if r2 != 0 {
		log.Println(r2, err, lpFreeBytesAvailable/1024/1024)
	}
}
