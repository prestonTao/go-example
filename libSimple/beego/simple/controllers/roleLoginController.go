package controllers

import (
	// "github.com/astaxie/beego"
	"log"
)

type RoleLoginGridx struct {
	GridxPageination
}

func (this *RoleLoginGridx) Get() {
	log.Println("haha")
	users := make([]User, 0)
	u1 := User{Id: 1, Title: "tao", Artist: "haha"}
	users = append(users, u1)
	u2 := User{Id: 2, Title: "tao", Artist: "haha"}
	users = append(users, u2)
	u3 := User{Id: 3, Title: "tao", Artist: "haha"}
	users = append(users, u3)
	u4 := User{Id: 4, Title: "tao", Artist: "haha"}
	users = append(users, u4)
	u5 := User{Id: 5, Title: "tao", Artist: "haha"}
	users = append(users, u5)
	u6 := User{Id: 6, Title: "tao", Artist: "haha"}
	users = append(users, u6)
	u7 := User{Id: 7, Title: "tao", Artist: "haha"}
	users = append(users, u7)
	u8 := User{Id: 8, Title: "tao", Artist: "haha"}
	users = append(users, u8)
	u9 := User{Id: 9, Title: "tao", Artist: "haha"}
	users = append(users, u9)
	u10 := User{Id: 10, Title: "tao", Artist: "haha"}
	users = append(users, u10)

	this.GetRange()
	this.SetTotal(100)

	this.Data["json"] = users
	this.ServeJson()
	this.StopRun()
}
