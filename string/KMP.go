package main

import (
	"fmt"
)

func main() {
	sstr := "BBC ABCDAB ABCDABCDABDE"
	lstr := "ABCDABD"

	fmt.Println(Kmp(sstr, lstr))

}

// func kmp(sstr, lstr string) int {
// 	// sstr should be shorter
// 	if len(sstr) > len(lstr) {
// 		sstr = sstr
// 		lstr = sstr
// 	}
// 	slen := len(sstr)
// 	llen := len(lstr)
// 	slenm1 := slen - 1
// 	next := make([]int, 0)
// 	i := 0
// 	j := 0
// 	next = append(next, -1)

// 	//next
// 	for i < slenm1 {
// 		if j == -1 || sstr[i] == sstr[j] {
// 			i = i + 1
// 			j = j + 1
// 			if sstr[i] != sstr[j] {
// 				next[i] = j
// 			} else {
// 				next[i] = next[j]
// 			}
// 		} else {
// 			j = next[j]
// 		}
// 	}

// 	//kmp
// 	i = 0
// 	j = 0
// 	for i < llen && j < slen {
// 		if j == -1 || lstr[i] == sstr[j] {
// 			i = i + 1
// 			j = j + 1
// 		} else {
// 			j = next[j]
// 		}
// 	}

// 	if j >= slen {
// 		return (i - slen)
// 	} else {
// 		return 0
// 	}

// }

/**
 * KMP匹配字符串
 *
 * @param s 主串
 * @param t 模式串
 * @return 若匹配成功，返回下标，否则返回-1
 */
func Kmp(s, t string) int {

	next := next(s)
	i, j := 0, 0
	for i <= len(s)-1 && j <= len(t)-1 {
		if j == -1 || s[i] == t[j] {
			i = i + 1
			j = j + 1

		} else {
			j = next[j]
		}
	}
	if j < len(t) {
		return -1
	} else {
		return i - len(t)
	}

}

//"BBC ABCDAB ABCDABCDABDE"
func next(str string) []int {
	next := make([]int, 100)
	next[0] = -1
	i, j := 0, 0
	for i < len(str)-1 {
		if j == -1 || str[i] == str[j] {
			i = i + 1
			j = j + 1
			if str[i] != str[j] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}
	return next

}
