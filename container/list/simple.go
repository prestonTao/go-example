package main

import (
	"container/list"
	"fmt"
)

func main() {
	simple1()
}

func simple1() {
	names := list.New()
	//插入一个元素
	names.PushFront("nihao")
	fmt.Println("集合长度为：", names.Len())
	//从集合中获得一个元素
	name := names.Back()
	fmt.Println("得到的元素：", name.Value)
	//从集合中删除一个元素
	names.Remove(name)
	fmt.Println("现在集合的长度为：", names.Len())
}
