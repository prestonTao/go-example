package main

import (
	"github.com/astaxie/beego"
	_ "libSimple/beego/simple/routers"
)

func main() {
	// beego.DirectoryIndex=true
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
