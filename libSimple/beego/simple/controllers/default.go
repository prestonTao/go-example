package controllers

import (
	"github.com/astaxie/beego"
	"log"
	// "strconv"
	// "strings"
)

// type MainController struct {
// 	beego.Controller
// }

// func (this *MainController) Get() {
// 	this.Data["Website"] = "beego.me"
// 	this.Data["Email"] = "astaxie@gmail.com"
// 	this.TplNames = "index.tpl"
// }

type HomeController struct {
	beego.Controller
}

//首页
func (this *HomeController) Get() {
	// log.Println("haha")
	// this.Ctx.Redirect(302, "/static/index")
	this.TplNames = "index.tpl"
}

func (this *HomeController) PageOne() {
	pagePath := this.Ctx.Input.Param(":pagePath")
	this.TplNames = "html/" + pagePath
}

type GridxDemo struct {
	GridxPageination
}

func (this *GridxDemo) Get() {
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

// func (this *GridxPageination) Get() {}
// func (this *GridxPageination) Get() {}
// func (this *GridxPageination) Get() {}
// func (this *GridxPageination) Get() {}
// func (this *GridxPageination) Get() {}
// func (this *GridxPageination) Get() {}

type PageOne struct {
	totalCount int
	items      []User
}

type User struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
