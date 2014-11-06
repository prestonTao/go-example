package main

import (
	"fmt"
	chash "github.com/golang/groupcache/consistenthash"
)

func main() {
	simple1()
	fmt.Println("-------------------------")
	simple2()
}

func simple1() {
	h := chash.New(3, nil)
	h.Add("2", "4")
	v := h.Get("6")
	fmt.Println(v)
}

func simple2() {
	h := chash.New(40, nil)
	h.Add("tao", "taop", "taopop")
	fmt.Println(h.Get("tao"))
}
