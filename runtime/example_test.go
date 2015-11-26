package runtime

import (
	"fmt"
	"runtime"
	"testing"
)

/*
	判断操作系统类型
*/
func TestCheckOS(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		fmt.Println("这是windows系统")
	}
	fmt.Println(runtime.GOARCH)
}
