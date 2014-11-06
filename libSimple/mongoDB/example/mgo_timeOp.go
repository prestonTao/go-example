package main

import (
	"../mgo"
	"../mgo/bson"
	"fmt"
	"time"
)

func main() {
	// str_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	// fmt.Println(str_time)
	insertTime()
	// selectTime()
}

func insertTime() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("oss").C("roleLogin")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639", time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)},
		//插入无时区时间格式
		&Person{"Cla", "+55 53 8402 8510", time.Date(2014, 1, 7, 5, 50, 4, 0, time.UTC)})
	if err != nil {
		panic(err)
	}
}

func selectTime() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("oss").C("roleLogin")

	//------------------------------
	//  第一种查询多条记录方式
	//------------------------------
	result := Person{}
	iter := c.Find(bson.M{"time": bson.M{"$gte": time.Date(2014, 1, 7, 0, 0, 0, 0, time.Local)}}).Iter()
	// iter := c.Find(bson.M{"name": "Ale"}).Iter()
	fmt.Println(iter)
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result)
		// personAll.Persons = append(personAll.Persons, result)
	}

	//------------------------------
	//  第二种查询多条记录方式
	//------------------------------
	resultAll := []Person{}
	err = c.Find(bson.M{"name": "Ale"}).All(&resultAll)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultAll)

}

type Person struct {
	Name  string
	Phone string
	Time  time.Time
}
