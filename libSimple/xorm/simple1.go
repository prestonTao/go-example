package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	simple1()
	simple2()
}

//创建表格
func simple1() {
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")

	defer engine.Close()
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	//设置显示sql语句
	engine.ShowSQL = true
	//设置最大连接数为5
	engine.SetMaxOpenConns(5)
	engine.CreateTables(&User{}, &Role{})
}

//插入数据
func simple2() {
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	defer engine.Close()
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	id, _ := engine.Insert(&User{Name: "tao", Age: 30})
	fmt.Println(id)

}

//执行一条sql语句，查询数据
func simple3() {
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	defer engine.Close()
	if err != nil {
		fmt.Println("连接数据库失败")
	}

	users := make([]User, 0)
	err = engine.Sql("select * from user where name='tao'").Find(&users)
	if err != nil {
		fmt.Println("查询失败")
	}
	fmt.Println(users)
}

type User struct {
	Id   uint64 `xorm:"pk autoincr 'c_id'" json:"c_id"` //id
	Name string
	Age  int32
}

type Role struct {
	User     `xorm:"extends"` //创建的表中包含User中的字段
	RoleName string
}
