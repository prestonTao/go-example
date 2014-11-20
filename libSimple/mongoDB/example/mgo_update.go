package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"mongo"
)

type Machine struct {
	Id_     bson.ObjectId `bson:"_id"`
	Frameid string        `bson:"frameid"`
	Ip      string        `bson:"ip"`
	Name    string        `bson:"name"`
	SType   string        `bson:"type"`
	SOS     string        `bson:"OS"`
	System  string        `bson:"system"`
	Effect  string        `bson:"effect"`
	Usage   string        `bson:"usage"`
}

type Qe struct {
	Id_   bson.ObjectId `bson:"_id"`
	Usage string
}
type result interface {
}

func main() {
	mm := mongo.NewMongoDBConn()
	mm.Connect("mongodb://redis:redis@localhost:27017/rrest")
	defer mm.Stop()
	mm.SetDBName("rrest")
	mm.SetTableName("machinea")
	selector := bson.M{"_id": bson.ObjectIdHex("527251bfda4f191e797258d4")}
	change := bson.M{"$set": bson.M{"usage": "hello world"}}
	mm.Update(selector, change)
	fmt.Printf("end\n")
}

/*
func main() {

	mm := mongo.NewMongoDBConn()
	mm.Connect("mongodb://redis:redis@localhost:27017/rrest")
	defer mm.Stop()
	mm.SetDBName("rrest")
	mm.SetTableName("machinea")
	mm.Insert(&Qe{
		Id_: bson.NewObjectId(),
		Usage:"eeeeea",
	})
    fmt.Printf("end\n")
} */
