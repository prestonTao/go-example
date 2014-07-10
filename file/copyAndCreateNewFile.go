package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	foo()
}

func foo() {
	oldFile := "file/xxx.json"
	newFile := "file/newfile.json"
	//file, err := os.Open(oldFile)
	bytes, err := ioutil.ReadFile(oldFile)
	if err != nil {
		fmt.Println("open file fiald  ||||  ", err)
	}
	//finfo, err1 := file.Stat()
	newfile, _ := os.Create(newFile)
	defer newfile.Close()
	newfile.Write(bytes)

}
