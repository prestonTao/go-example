package controllers

import (
	"github.com/astaxie/beego"
	// "log"
	"strconv"
	"strings"
)

type GridxPageination struct {
	beego.Controller
	startIndex int
	lastIndex  int
}

func (this *GridxPageination) GetRange() (int, int) {
	rangeValue := strings.Split(strings.Split(this.Ctx.Input.Header("Range"), "=")[1], "-")
	this.startIndex, _ = strconv.Atoi(rangeValue[0])
	this.lastIndex, _ = strconv.Atoi(rangeValue[1])
	return this.startIndex, this.lastIndex
}
func (this *GridxPageination) SetTotal(total int) {
	startIndexStr := strconv.Itoa(this.startIndex)
	lastIndexStr := strconv.Itoa(this.lastIndex)
	this.Ctx.Output.Header("Content-Range", "items="+startIndexStr+"-"+lastIndexStr+"/100")
}
