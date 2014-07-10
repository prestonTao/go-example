package inc

// #include "foo.h"
import "C"

func CallLib() {
	C.foo()
}
