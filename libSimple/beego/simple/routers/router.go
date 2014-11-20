package routers

import (
	"github.com/astaxie/beego"
	"libSimple/beego/simple/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/view/:pagePath", &controllers.HomeController{}, "get:PageOne")
	beego.Router("/gridxdemo/:all", &controllers.GridxDemo{})

	beego.Router("/roleLogingridx/:all", &controllers.RoleLoginGridx{})

}
