package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var GlobalPlayerDB *mgo.Collection //保存玩家信息
var GlobalCardDB *mgo.Collection   //保存充值卡信息

var GlobalPlayerManager *mgo.Collection    //保存玩家管理数据
var GlobalCoinChangeRecord *mgo.Collection //金币修改记录
var GlobalOrder *mgo.Collection            //订单

func init() {
	dbsession, err := mgo.Dial("127.0.0.1:27017/gd")
	if err != nil {
		panic(err)
	}
	dbsession.SetMode(mgo.Strong, true)
	db := dbsession.DB("gd")

	GlobalPlayerDB = db.C("player")
	GlobalCardDB = db.C("cards")
	GlobalPlayerManager = db.C("playerM")
	GlobalCoinChangeRecord = db.C("coinChangeRecord")
	GlobalOrder = db.C("order")
}

type PlayerManagerDB struct {
	PlayerCount int
}

func main() {
	var pmdbs []PlayerManagerDB
	GlobalPlayerManager.Find(nil).All(&pmdbs)
	fmt.Println(pmdbs)
}
