package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var UPLOAD_DIR = "D:/test/"

func main() {
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/file", File)
	http.ListenAndServe(":80", nil)
	// http.FileServer(root)
}

//文件上传
func Upload(w http.ResponseWriter, r *http.Request) {

}

func File(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	buf, _ := ioutil.ReadFile(UPLOAD_DIR + name)
	// w.Header().Set("Content-Type", "image")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	w.Header().Set("Content-Disposition", "attachment;filename="+name)
	w.Header().Set("Content-Type", "application")
	w.Write(buf)
	// http.ServeFile(w, r, imagePath)
}
