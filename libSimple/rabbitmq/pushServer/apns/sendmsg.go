package main

import (
	"encoding/json"
	"fmt"
	"github.com/virushuo/Go-Apns"
	"io/ioutil"
	"os"
	"time"
)

// type Config struct {
// 	ApnsCert string `json:"apns_cert"`
// 	ApnsKey  string `json:"apns_key"`
// }

func loadConfig() map[string]interface{} {
	bytes, err := ioutil.ReadFile("conf/development.json")
	if err != nil {
		fmt.Println("读配置文件路径错误：", err.Error())
	}

	var dat map[string]interface{}
	// config := new(Config)
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}
	return dat
}

func main() {

	configs := loadConfig()
	apnsCert, _ := configs["apns_cert"].(string)
	apnsKey, _ := configs["apns_key"].(string)

	apn, err := apns.New(apnsCert, apnsKey, "gateway.sandbox.push.apple.com:2195", 1*time.Second)
	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		os.Exit(1)
	}
	go readError(apn.ErrorChan)

	token := "710617f14b4abfe59a9ba6c0ef7ffeeabe31797ebb7ec183d342efba6c9295a7"

	payload := apns.Payload{}
	payload.Aps.Alert.Body = "hello world! 0"

	notification := apns.Notification{}
	notification.DeviceToken = token
	notification.Identifier = 0
	notification.Payload = &payload
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)

	notification.Identifier++
	notification.Payload.Aps.Alert.Body = "hello world! 1"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)

	notification.Identifier++
	notification.Payload.Aps.Alert.Body = "hello world! 2"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)

	notification.Identifier++
	notification.DeviceToken = ""
	notification.Payload.Aps.Alert.Body = "hello world! 3"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)
	time.Sleep(1E9)

	notification.Identifier++
	notification.DeviceToken = token
	notification.Payload.Aps.Alert.Body = "re hello world! 0"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)

	notification.Identifier++
	notification.DeviceToken = ""
	notification.Payload.Aps.Alert.Body = "re hello world! 1"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)
	time.Sleep(1E9)

	notification.Identifier++
	notification.DeviceToken = token
	notification.Payload.Aps.Alert.Body = "rere hello world! 0"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)
	time.Sleep(2e9)

	notification.Identifier++
	notification.DeviceToken = token
	notification.Payload.Aps.Alert.Body = "rere hello world! 1"
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)
	time.Sleep(2e9)

	apn.Close()
}

func readError(errorChan <-chan error) {
	for {
		apnerror := <-errorChan
		fmt.Println(apnerror.Error())
	}
}
