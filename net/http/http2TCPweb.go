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
	http.HandleFunc("/src.js", srcjs)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
	fmt.Println("webServer startup...")
}

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := r.URL.Path
	scheme := r.URL.Scheme
	log.Println(path, "     ", scheme)

	pagesrc := `<html>
		<head>
			<title>index</title>
			<script src="src.js"></script>
		</head>
		<body>
			index
		</body>
		</html>`

	io.WriteString(w, pagesrc)
}

func nimei(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := r.URL.Path
	scheme := r.URL.Scheme
	log.Println(path, "     ", scheme)
	io.WriteString(w, "nimei")
}

func srcjs(w http.ResponseWriter, r *http.Request) {
	log.Println("srcjs")
	io.WriteString(w, "alert(1);")
}
