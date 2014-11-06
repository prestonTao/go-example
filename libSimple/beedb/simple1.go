package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	simple1()
}

func simple1() {
	db, err := sql.Open("mymysql", "test/root/root")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)

	var u User
	orm.Where("id = 1").Find(u)
	fmt.Println(u)

	// rac, err := orm.Exec("select * from t_user")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(rac.RowsAffected())
}

type User struct {
	Id   int32
	Name string
	Age  int32
}
