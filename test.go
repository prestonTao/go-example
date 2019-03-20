package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.FormatFloat(1.23456, 'f', 3, 32)
	fmt.Println(a)

}
