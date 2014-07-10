package include

// // some comment
//#include "inc.h"
import "C"

func Output(s string) {
	p := C.CString(s)
	C.output(p)
}

func OutTwo(s string) {
	p := C.CString(s)
	C.outTwo(p)
}

func OutThree(s string) {
	p := C.CString(s)
	C.outThree(p)
}
