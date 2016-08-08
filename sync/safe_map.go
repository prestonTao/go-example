package main

import (
	"fmt"
	"sync"
)

func main() {
	one := mem[1]
	delete(mem, 1)
	one = one + 1
	fmt.Println(mem)
}

var mem = map[int]int{1: 1, 2: 2}
var lock = new(sync.RWMutex)

func AddUpdate(key, value int) {
	lock.Lock()
	mem[key] = value
	lock.Unlock()
}

func Delete(key int) {
	delete(mem, key)
}
func Find(key int) int {
	lock.RLock()
	one := mem[key]
	lock.RUnlock()
	return one
}
