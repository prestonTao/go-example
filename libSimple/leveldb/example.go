package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var db *leveldb.DB

func main() {
	inition()
	//	example()
	//	insert()
	//	find()
	//	insertImg()

	findImg("d7dda8c95b8f1a614f0aa6871fe4724b",
		"3329c9528aa5e1a46f44f420a57a7829",
		"aa9713ce9fb9656d7793b26a844950c5",
		"572a74957978db25a55903550eb3941b",
		"9d750694afb9f70202ba49b0aa9c56ab",
		"69e4e3fa1e2474fe2ecfb8341caad932",
		"1bbf3cbe992e4514ca2699e539dd1d37",
		"52cefa98bea1a445edf1e03b9d8a487f")

	//	insertLittle()

}

func inition() {
	//没有db目录会自动创建
	var err error
	db, err = leveldb.OpenFile("db", nil)
	//	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func example() {

	err := db.Put([]byte("tao"), []byte("陶"), nil)
	if err != nil {
		fmt.Println("111", err)
	}
	value, err := db.Get([]byte("tao"), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value", string(value))
}

func insert() {
	for i := 0; i < 100; i++ {
		err := db.Put([]byte("name_"+strconv.Itoa(i)), []byte(strconv.Itoa(i)), nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < 100; i++ {
		err := db.Put([]byte("age_"+strconv.Itoa(i)), []byte(strconv.Itoa(i)), nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func find() {
	iter := db.NewIterator(util.BytesPrefix([]byte("name_")), nil)
	for iter.Next() {
		fmt.Println("key:", string(iter.Key()), "value:", string(iter.Value()))
	}
	iter.Release()
	err := iter.Error()
	fmt.Println(err)
}

func insertImg() {
	bs, err := ioutil.ReadFile("C:/Users/preston/Desktop/image/图片/7af40ad162d9f2d3de7180f7a5ec8a136327cc42.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 1024*3; i++ {
		h := md5.New()
		h.Write([]byte(time.Now().String() + strconv.Itoa(i)))
		name := h.Sum(nil)
		fmt.Println(hex.EncodeToString(name))

		t1 := time.Now()
		err := db.Put([]byte(name), bs, nil)

		if err != nil {
			fmt.Println("111", err)
		}
		fmt.Println(i, hex.EncodeToString(name), time.Now().Sub(t1)) //计算出微毫秒，十亿份之一秒
	}
}

func findImg(names ...string) {
	for _, one := range names {
		key, err := hex.DecodeString(one)
		if err != nil {
			return
		}
		t1 := time.Now()
		value, err := db.Get(key, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("value", one, time.Now().Sub(t1), len(value))
	}

}

func insertLittle() {

	for i := 0; i < 1024*3; i++ {
		h := md5.New()
		h.Write([]byte(time.Now().String() + strconv.Itoa(i)))
		name := h.Sum(nil)

		t1 := time.Now()
		err := db.Put(name, name, nil)

		if err != nil {
			fmt.Println("111", err)
		}
		fmt.Println(i, hex.EncodeToString(name), time.Now().Sub(t1).Nanoseconds()) //计算出微毫秒，十亿份之一秒
	}
}

func findLittle(names ...string) {
	for _, one := range names {
		key, err := hex.DecodeString(one)
		if err != nil {
			return
		}
		t1 := time.Now()
		value, err := db.Get(key, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("value", one, time.Now().Sub(t1), len(value))
	}

}
