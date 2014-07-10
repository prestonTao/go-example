package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("f", "Null", "please input a file name or dir.")
var size *string = flag.String("s", "0", "please input a dst file size.")

func SplitFile(file *os.File, size int, target string) {
	finfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed:", file, size)
		return
	}

	//每次最多拷贝1m
	bufsize := 1024 * 1024
	if size < bufsize {
		bufsize = size
	}

	buf := make([]byte, bufsize)
	num := (int(finfo.Size()) + size - 1) / size
	for i := 0; i < num; i++ {
		copylen := 0
		newfilename := target + finfo.Name() + strconv.Itoa(i)
		newfile, err1 := os.Create(newfilename)
		if err1 != nil {
			fmt.Println("failed to create file", newfilename)
			defer newfile.Close()
			continue
		} else {
			fmt.Println("create file:", newfilename)
		}

		for copylen < size {
			n, err2 := file.Read(buf)
			if err2 != nil && err2 != io.EOF {
				fmt.Println(err2, "failed to read from:", file)
				break
			}

			if n <= 0 {
				break
			}

			//写文件
			w_buf := buf[:n]
			newfile.Write(w_buf)
			defer newfile.Close()
			copylen += n
		}
	}
}

func main() {
	flag.Parse()

	if *infile == "Null" {
		fmt.Println("no file to input")
		return
	}

	file, err := os.Open(*infile)
	if err != nil {
		fmt.Println("failed to open:", *infile)
		defer file.Close()
		return
	}

	size, _ := strconv.Atoi(*size)

	SplitFile(file, size*1024, "C:\\")

	defer file.Close()
}
