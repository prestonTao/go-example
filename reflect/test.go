// ConstomRouter project main.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	test1()
}

func test1() {
	uType := reflect.TypeOf(User{})
	uValue := reflect.New(uType)
	fmt.Println(uValue)
	fmt.Println(uValue.Elem().CanSet())
	fmt.Println(uValue.Elem().CanAddr())
	fieldName := uValue.Elem().FieldByName("name").Addr()
	fmt.Println(fieldName)

	//a := User{}
	//b := new(User)
	//fmt.Println(a, "--", b)
	//valueA := reflect.ValueOf(a)
	//valueB := reflect.ValueOf(b)
	//fmt.Println(valueA, "--", valueB)
	//nameValueA := valueA.FieldByName("name")
	//fmt.Println(nameValueA.)
	////nameValueB := valueB.FieldByName("name")
	//fmt.Println(nameValueA.CanSet(), nameValueA.CanAddr())
	////fmt.Println(valueB.FieldByName("name"))
}

type User struct {
	name string
	age  int
}

func annotation(name string, pwd string) {
	log.Println("  ")
}
