package main

import (
	"fmt"
	"github.com/alyu/configparser"
	"time"
)

func main() {
	example2()
}

func example2() {
	for range time.NewTicker(time.Second).C {
		fmt.Println("nihao")
	}
}

func example1() {
	config, _ := configparser.Read("client.properties")
	section, err := config.Section("global")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(section)
}
