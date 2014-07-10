package memoryGC

//
import "C"

func MemoryMan() {
	//初始化gostr
	var gostr string
	cstr := C.CString(gostr)
	defer C.free(unsafe.Pointer(cstr))
	//  接下来大胆的使用cstr吧，因为可以保证被释放掉了
	//C.sprintf(cstr)
}
