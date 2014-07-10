package main

import (
	"hover"
	// "include"
	"inc"
	"prints"
)

func main() {
	// runInclude()
	runInc()
	runPrints()
	runHover()
}

// func runInclude() {
// 	include.Output("inc hello")
// 	include.OutTwo("inc hello Two")
// }
func runPrints() {
	prints.Prints("haha")
}

func runHover() {
	hover.Msgbox("title", "body")
}

func runInc() {
	inc.CallLib()
}
