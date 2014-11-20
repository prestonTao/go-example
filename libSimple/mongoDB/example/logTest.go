package main

import (
	"../mgo"
	"../mgo/bson"
	"encoding/json"
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

	var chg interface{}
	err = json.Unmarshal([]byte(`{"name":"tao","age":12}`), &chg)
	if err != nil {
		panic(err)
	}

	times := time.Now()
	err = c.Insert(chg, times)
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
