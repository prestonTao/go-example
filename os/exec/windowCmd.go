package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func init() {
}

func main() {
	test1()
	test2()
}

func test1() {
	cmd := exec.Command("ping", "baidu.com")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err.Error())
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

func test2() {
	cmd := exec.Command("ping", "baidu.com")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	fmt.Println("stdout: ", string(bytes))

}
