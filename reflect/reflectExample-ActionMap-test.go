// ConstomRouter project main.go
package main

import (
	"fmt"
	"reflect"
)

type Action func(w, r string)

type __ActionMap struct {
	actionMap map[string]Action
}

func NewAction(actionName string, fun Action) {
	if ActionMap.actionMap == nil {
		ActionMap.actionMap = make(map[string]Action)
	}
	ActionMap.actionMap[actionName] = fun
}

func GetAction(actionName string) Action {
	if k, v := ActionMap.actionMap[actionName]; v {
		return k
	}
	return nil
}

var ActionMap = __ActionMap{}

func main() {
	/*
	   var tempAction1=GetAction("action1")
	   fmt.Println(tempAction1)
	   tempAction1("response1","request1")

	   var tempAction2=GetAction("action2")
	   fmt.Println(tempAction2)
	   tempAction2("response2","request2")
	*/
	var config_file_loaded = map[string]string{"/index/": "action1", "/auth/": "action2"}
	var routers = make(map[string]Action)
	for url, action_name := range config_file_loaded {
		var tempAction = GetAction(action_name)
		routers[url] = tempAction
		tempAction(url, action_name)
	}
}

func init() {
	//NewAction("action1", func(w, r string) {
	//	fmt.Println(w, r)
	//})
	//NewAction("action2", func(w, r string) {
	//	fmt.Println("new ", w, r)
	//})
	s = new(HomeController)
	fun := reflect.ValueOf(a).MethodByName("GetName")
	typ.NumMethod()

	typ = typ.Elem()
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		NewAction("action1", p)
	}
}

type HomeController struct{}

func (c *HomeController) Home(w, r string) {
	fmt.Println(w, r)
}

func (c *HomeController) Hello(w, r string) {
	fmt.Println("new ", w, r)
}
