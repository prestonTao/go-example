package main

import (
	"fmt"
	"log"
	"time"

	"gitee.com/prestonTao/utils"
	"github.com/btcsuite/btcwallet/walletdb"
	_ "github.com/btcsuite/btcwallet/walletdb/bdb"
)

func main() {
	// exampleBdb()
	WriteAndRead()
}

func exampleBdb() {
	fmt.Println("start")

	spvdb, err := walletdb.Create("bdb", "test.db", true, time.Second*10)
	if err != nil {
		log.Printf("Unable to create Neutrino DB: %s", err)
		return
	}
	spvdb.Close()

	spvdb, err = walletdb.Open("bdb", "test.db", true, time.Second*10)
	if err != nil {
		log.Printf("Unable to open Neutrino DB: %s", err)
		return
	}
	defer spvdb.Close()

	key1 := []byte("key1")
	key2 := []byte("key2")

	key := []byte("key")
	value := []byte("value")

	err = walletdb.Update(spvdb, func(tx walletdb.ReadWriteTx) error {
		ns, err := tx.CreateTopLevelBucket(key1)
		if err != nil {
			return err
		}

		bucket, err := ns.CreateBucketIfNotExists(key2)
		if err != nil {
			return err
		}
		err = bucket.Put(key, value)
		if err != nil {
			fmt.Println("保存错误:", err.Error())
			return nil
		}
		return nil
	})
	if err != nil {
		fmt.Println("保存错误", err.Error())
		return
	}

	err = walletdb.View(spvdb, func(tx walletdb.ReadTx) error {
		ns := tx.ReadBucket(key1)

		bucket := ns.NestedReadBucket(key2)

		value := bucket.Get(key)
		fmt.Println("查询到的value:", string(value))

		return nil
	})

	if err != nil {
		fmt.Println("查询错误", err.Error())
		return
	}

}

/*
读写性能测试
*/
func WriteAndRead() {
	fmt.Println("start")

	spvdb, err := walletdb.Create("bdb", "test.db", true, time.Second*10)
	if err != nil {
		log.Printf("Unable to create Neutrino DB: %s", err)
		return
	}

	Write(spvdb)
	Read(spvdb)
	spvdb.Close()
}

/*
写数据库性能测试
*/
func Write(db walletdb.DB) {
	start := time.Now()

	for i := 0; i < 10000; i++ {
		key1 := utils.Uint64ToBytesByBigEndian(uint64(i))
		// key1 := []byte("key1")
		// key2 := []byte("key2")

		key := []byte("key")
		value := []byte("value")

		err := walletdb.Update(db, func(tx walletdb.ReadWriteTx) error {
			ns, err := tx.CreateTopLevelBucket(key1)
			if err != nil {
				return err
			}

			// bucket, err := ns.CreateBucketIfNotExists(key2)
			// if err != nil {
			// 	return err
			// }
			// err = bucket.Put(key, value)
			err = ns.Put(key, value)
			if err != nil {
				fmt.Println("保存错误:", err.Error())
				return nil
			}
			return nil
		})
		if err != nil {
			fmt.Println("保存错误", err.Error())
			return
		}
	}

	fmt.Println("写数据库耗时", time.Now().Sub(start))
}

/*
读数据库性能测试
*/
func Read(db walletdb.DB) {
	start := time.Now()

	fmt.Println("读数据库耗时", time.Now().Sub(start))
}
