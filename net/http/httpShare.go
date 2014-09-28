package main

import (
	"fmt"
	"net/http"
)

//文件服务器
func main() {
	h := http.FileServer(http.Dir("."))
	var port string
	fmt.Printf("请输入端口号: ")
	fmt.Scanf("%s", &port)
	http.ListenAndServe(":"+port, h)
}
