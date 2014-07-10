// ConstomRouter project main.go
package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		hello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello myrouter")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":80", mux)
}
