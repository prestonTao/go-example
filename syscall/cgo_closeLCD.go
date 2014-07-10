//功能: 关闭屏幕显示器,节约电费
//作者: 维尼熊
//参考: http://www.dreamincode.net/forums/topic/148684-turn-off-your-lcd/
package main

/*
#include "windows.h"
void run(){
	SendMessageA(FindWindowA(0,0),0x112,0xf170,2);
}
*/
import "C"

func main() {
	C.run()
}
