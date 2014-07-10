package main

import (
	//"bytes"
	"fmt"
	// "io/ioutil"
	"os/exec"
)

func main() {
	test1()
}

func test1() {
	cmd := exec.Command("ping", "baidu.com")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stdout))
}
