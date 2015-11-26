package main

import (
	"fmt"
)

func main() {
	fmt.Println(show_strlen("哈哈1"))
	fmt.Println(show_substr("哈哈1什1么的", 9))
}
func show_substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

/*
	计算字符串长度，中文占2个长度，英文占1个长度
*/
func show_strlen(s string) int {
	sl := 0
	rs := []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}
