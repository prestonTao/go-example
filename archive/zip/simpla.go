package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	simple1()
}

func simple1() {
	b := bytes.NewBuffer([]byte("haha"))
	fmt.Println(b.Bytes())
	z := zip.NewWriter(b)
	defer z.Close()

	fmt.Println(b.Bytes())

	// zip.Writer([]byte())
}
