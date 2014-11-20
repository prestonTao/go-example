// ConstomRouter project main.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		hello(w, r)
		return
	}
	if r.URL.Path == "/nimei" {
		nimei(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello start")
	time.Sleep(time.Second * 10)
	fmt.Fprintf(w, "hello myrouter")
	fmt.Println("hello end")
}

func nimei(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nimei start")
	time.Sleep(time.Second * 10)
	fmt.Fprintf(w, "nimei myrouter")
	fmt.Println("nimei end")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":80", mux)
}
