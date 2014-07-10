package prints

// // some comment
//#include "prints.h"
import "C"

func Prints(s string) {
	p := C.CString(s)
	C.prints(p)
}
