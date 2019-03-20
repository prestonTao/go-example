package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := new(sync.WaitGroup)
	go func() {
		time.Sleep(1 * time.Second)
		group.Add(1)
		group.Done()
		group.Done()
		group.Done()
		group.Done()
	}()
	group.Wait()
	fmt.Println("end")
}
