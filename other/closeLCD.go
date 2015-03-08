//功能: 关闭屏幕显示器,节约电费
//作者: 维尼熊
//参考: http://www.dreamincode.net/forums/topic/148684-turn-off-your-lcd/
package main
import "syscall"
func main() {
	u := syscall.MustLoadDLL("user32.dll")
	win, _, _ := u.MustFindProc("FindWindowA").Call(0, 0)
	u.MustFindProc("SendMessageA").Call(win, 0x112, 0xf170, 2)
}