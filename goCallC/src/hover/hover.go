package hover

//#include <windows.h>
import "C"

func Msgbox(title string, body string) int {
	C.MessageBox(nil, (*C.CHAR)(C.CString(body)), (*C.CHAR)(C.CString(title)), 0)
	return 0
}
