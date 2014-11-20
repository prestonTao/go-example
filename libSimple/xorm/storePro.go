package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	// engine, err := xorm.NewEngine("mymysql", "root:root@/test?charset=utf8")
	// engine, err := xorm.NewEngine("mymysql", "test/root/root*127.0.0.1,tcp:127.0.0.1,laddr=127.0.0.1,timeout=23")

	engine, err := xorm.NewEngine("mymysql", "test/root/root*127.0.0.1,tcp:127.0.0.1,laddr=127.0.0.1,timeout=23")

	defer engine.Close()
	if err != nil {
		fmt.Println("连接数据库失败", err.Error())
	}
	//设置显示sql语句
	engine.ShowSQL = true
	//设置最大连接数为5
	engine.SetMaxOpenConns(5)
	r, e := engine.Exec("call tao(1)")
	if e != nil {
		fmt.Println("存储过程调用失败")
	}
	fmt.Println(r)
}

type User struct {
	Id   int32
	Name string
	Age  int32
}
