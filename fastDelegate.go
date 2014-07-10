package main

import(
"fmt"
"reflect"
)

func main() {
}


type Event struct{
	//要执行方法的对象
	Obj interface{}
	//要执行方法的名称
	methodName string
	//要执行方法的类型
	Params []string

}

//根据参数数组生成参数类型数组
func (this *Event) contractParamTypes(Object[] params){

}

//执行该对象的该方法
func (this *Event) invoke(){

}



//EventHandler类，若干Event类的载体，同时提供一个执行所有Event的方法
type EventHandler struct{
	enevt []Event
}

func (this *EventHandler) addEvent(enent Event){

}
