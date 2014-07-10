package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/nimei", nimei)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
	fmt.Println("webServer startup...")
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello World")
}

func nimei(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "nimei")
}
