package main

// import (
// 	"../mgo"
// 	"../mgo/bson"
// 	"fmt"
// )

// func main() {
// 	find()
// }

// type Person struct {
// 	Name  string
// 	Phone string
// }

// func insert() {
// 	//session, err := mgo.Dial("server1.example.com,server2.example.com")
// 	session, err := mgo.Dial("127.0.0.1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("test").C("people")
// 	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
// 		&Person{"Cla", "+55 53 8402 8510"})
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func find() {
// 	//session, err := mgo.Dial("server1.example.com,server2.example.com")
// 	session, err := mgo.Dial("127.0.0.1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)

// 	c := session.DB("test").C("people")

// 	result := Person{}
// 	err = c.Find(bson.M{"name": "Ale"}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Phone:", result.Phone)
// }
