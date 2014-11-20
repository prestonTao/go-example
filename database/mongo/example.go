package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var FileDBName = "fs"           //存放二进制文件集合
var NameCollection = "name"     //存放二进制文件名和真实文件名对应关系的集合
var ChannelInfo = "channelinfo" //存放渠道信息集合

var MDB *mgo.Database

func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic("mgo init errors")
	}
	session.SetMode(mgo.Monotonic, true)
	MDB = session.DB("fileServer")
}

func find() {
	var info map[string]string = make(map[string]string)
	MDB.C(ChannelInfo).Find(bson.M{"channeluserid": "2"}).One(&info)
	delete(info, "channeluserid")
	delete(info, "_id")
	for key, value := range info {
		fmt.Println(key, value)
	}

}

func main() {
	find()
}
