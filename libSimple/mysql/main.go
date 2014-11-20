// mysql project main.go
package main

import (
	_ "Go-SQL-Driver/mysql-master"
	"database/sql"
	"fmt"

//"time"
)

func main() {
	//insertDB()
	//updateDB()
	//selectDB()
	delDB()
}

func insertDB() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	db.Close()
}

func updateDB() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	id := 1
	res, err := stmt.Exec("astaxieupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func selectDB() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	db.Close()
}

func delDB() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	id := 1
	res, err := stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
