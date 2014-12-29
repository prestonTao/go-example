package main

import (
	"fmt"
	"github.com/alyu/configparser"
)

func main() {
	config, _ := configparser.Read("client.properties")
	section, err := config.Section("global")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(section)
}
