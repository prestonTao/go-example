package main

import (
	// "../mgo"
	// "../mgo/bson"
	"fmt"
)

func main() {
	find()
}

func find() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("users")

	result := User{}
	err = c.Find(bson.M{"Name": "陈萌"}).One(&result)
	//err = c.Find(bson.M{"Name": "Name"}).One(&result)
	//result := []User{}
	//err = c.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

//func findTest(){

//}

type User struct {
	_id       string
	Name      string
	CardNo    string
	Descriot  string
	CtfTp     string
	CtfId     string
	Gender    string
	Birthday  int64 //生日
	Address   string
	Zip       int
	Dirty     string
	District1 string
	District2 string //可能是年龄
	District3 int
	District4 int
	District5 string
	District6 string
	FirstNm   string
	LastNm    string
	Duty      string
	Mobile    int    //电话：携程网147541773担保
	Tel       string //座机？
	Fax       string
	EMail     string
	Nation    string //名族
	Taste     string
	Education string
	Company   string
	CTel      string
	CAddress  string
	CZip      string
	Family    int
	Version   string //时间
	id        int    //id
}
