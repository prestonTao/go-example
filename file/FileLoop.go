// FileLoop
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("d://test", fileDir)
}

var fileMap = make(map[string]string)

func fileDir(path string, f os.FileInfo, err error) error {
	//if f == nil {
	//	return err
	//}
	//if f.IsDir() {
	//	return nil
	//}
	fmt.Println(path)
	return nil
}
