package main

import (
	"example/libSimple/alidayu/utils"
	"fmt"
	"net"
	"time"
)

func main() {
	SendToo()
}

func Start() {
	fmt.Println("start")
	done := false
	for !done {
		conn, err := net.Dial("tcp4", "120.76.188.148:8090")
		if err != nil {
			Send()
			done = true
			break
		}
		conn.Close()
		time.Sleep(time.Minute * 1)
	}
	fmt.Println("end")
}

func Send() {
	utils.AppKey = "23312050"
	utils.AppSecret = "db90086690a148669f7602496d24e956"

	//	success, resp := utils.SendSMS("13408066190,13551322482", "身份验证", "SMS_5038149", `{"code":"1234","product":"阿里大鱼"}`)
	//	success, resp := utils.SendSMS("13408066190", "身份验证", "SMS_13236641", `{"code":"1234"}`)
	success, resp := utils.SendSMS("13408066190", "身份验证", "SMS_13272034", `{}`)
	if !success {
		fmt.Println("接口调用失败")
		return
	}
	code := utils.CheckSMSResult(resp)

	fmt.Println(code)
}

func SendToo() {
	utils.AppKey = "23312050"
	utils.AppSecret = "db90086690a148669f7602496d24e956"

	//	success, resp := utils.SendSMS("13408066190,13551322482", "身份验证", "SMS_5038149", `{"code":"1234","product":"阿里大鱼"}`)
	//	success, resp := utils.SendSMS("13408066190", "身份验证", "SMS_13236641", `{"code":"1234"}`)
	success, resp := utils.SendSMS("13408066190", "身份验证", "SMS_26145205", `{}`)
	if !success {
		fmt.Println("接口调用失败")
		return
	}
	code := utils.CheckSMSResult(resp)

	fmt.Println(code)
}
