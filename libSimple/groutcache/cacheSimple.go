package main

import (
	"fmt"
	"github.com/golang/groupcache/lru"
)

func main() {
	simple1()
}

func simple1() {
	//创建一个无限大的cache
	cache := lru.New(0)
	cache.Add("nihao", "wobuhao")
	v, _ := cache.Get("nihao")
	fmt.Println(v)
}
