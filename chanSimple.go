package main

import(
"fmt"
// "time"
"reflect"
)

func main() {
	c := make(chan int)
	 putChan(c)
	// time.Sleep(time.Second * 10)
	// content := <-c
	// fmt.Println(content)
}

func putChan(chans chan int){
	for chanone := range chans{
	 	fmt.Println(reflect.TypeOf(chanone))

	}
}