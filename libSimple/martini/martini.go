package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(func() {
		fmt.Println("执行了我的中间件")
	})

	m.Get("/", func() string {
		return "Hello world!"
	})

	//   m.Get("/", func() {
	//   println("hello world")
	// })

	// m.Get("/", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
	//   res.WriteHeader(200) // HTTP 200
	// })

	//
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	m.Get("/auth", Auth, func() string {
		// 该方法将会在authorize方法没有输出结果的时候执行.
		fmt.Println("这里应该不执行")
		return "nimei"
	})

	m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) string {
	fmt.Println("auth haha")
	return "nimei auth"
}
