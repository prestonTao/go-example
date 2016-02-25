package main

import (
	"fmt"
	"github.com/prestonTao/go-example/libSimple/alidayu/utils"
)

func main() {
	utils.AppKey = "23312050"
	utils.AppSecret = "db90086690a148669f7602496d24e956"

	success, resp := utils.SendSMS("13408066190,13551322482", "身份验证", "SMS_5038149", `{"code":"1234","product":"阿里大鱼"}`)
	if !success {
		fmt.Println("接口调用失败")
		return
	}
	code := utils.CheckSMSResult(resp)

	fmt.Println(code)
}
