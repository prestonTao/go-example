// ConstomRouter project main.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	test3()
}

func test1() {
	// iterate through the attributes of a Data Model instance
	for name, mtype := range attributes(&Dish{}) {
		fmt.Printf("Name: %s, Type %s\n", name, mtype.Name())
	}
}

type Dish struct {
	Id     int
	Name   string
	Origin string
	Query  func()
}

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

//------------------------------------------------

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName(name string) string {
	fmt.Println(name)
	return this.name
}

func test2() {
	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(s))
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)

	fmt.Println(typ.NumMethod())
	m := typ.Method(0)
	fmt.Println(m.Type, "========", m.PkgPath, "======", m.Func)

	fmt.Println("-------------------")
	//方法的传入参数
	fmt.Println("方法中的参数总数：", m.Type.NumIn())
	fmt.Println("方法参数中是否有“...”类型参数：", m.Type.IsVariadic())
	fmt.Println("第一个参数的类型：", m.Type.In(0), "第二个参数的类型：", m.Type.In(1))
	//方法的返回值
	fmt.Println("方法返回值总数：", m.Type.NumOut())
	fmt.Println("第一个返回值:", m.Type.Out(0))

	fmt.Println("-------------------")

	str := "test"
	fun := reflect.ValueOf(a).MethodByName("GetName")
	b := fun.Call([]reflect.Value{reflect.ValueOf(str)})
	fmt.Println(b[0])

	//aa := reflect.TypeOf(a).MethodByName("GetName")
	//fmt.Println(fun.Name)
	fmt.Println("000")
}

//-------------------------------------------------
func test3() {
	fmt.Println("--------------")
	var a MyStruct
	b := new(MyStruct)
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b))

	fmt.Println("--------------")
	a.name = "yejianfeng"
	b.name = "yejianfeng"
	val := reflect.ValueOf(a).FieldByName("name")

	//painc: val := reflect.ValueOf(b).FieldByName("name")
	fmt.Println(val)

	fmt.Println("--------------")
	fmt.Println(reflect.ValueOf(a).FieldByName("name").CanSet())
	fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

	fmt.Println("--------------")
	var c string = "yejianfeng"
	p := reflect.ValueOf(&c)
	fmt.Println(p.CanSet())        //false
	fmt.Println(p.Elem().CanSet()) //true
	p.Elem().SetString("newName")
	fmt.Println(c)
}
