package main

/*
#include <stdio.h>
void hello(){
	printf("Hello cgo2");
}
*/
import "C"

func Hello() {
	C.hello()
}

func main() {
	Hello()
}
