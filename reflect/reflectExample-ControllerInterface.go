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
	//c := new(HomeController)

	//c.Hello("t", "ao")
	controllerMap["HomeController"] = &HomeController{}
	c := controllerMap["HomeController"]
	typ := reflect.TypeOf(c)
	fmt.Println(typ.NumMethod())
	//得到方法
	for i := 0; i < 2; i++ {
		fun := typ.Method(i)
		//fun := typ.MethodByName("Home")
		//查询方法参数个数
		fmt.Println(fun.Type.NumIn())
		fmt.Println(fun.Name)
		if fun.Name == "Home" {
			result := reflect.ValueOf(c).MethodByName("Home").Call([]reflect.Value{reflect.ValueOf("t"), reflect.ValueOf("ao")})
			fmt.Println(result)
		}
	}

	//fun := controllerMap["c1"]
	//fun("ta", "o")
}

//type ControllerInterface func(w, r string) interface{}

type HomeController struct{}

func (c *HomeController) Home(w, r string) interface{} {
	fmt.Println(w, r)
	return "home"
}
func (c *HomeController) Hello(w, r string) interface{} {
	fmt.Println(w, r)
	return "hello"
}

var controllerMap = make(map[string]interface{})
