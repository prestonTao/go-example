package main

import (
	// "bufio"
	"fmt"
	// "os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "baidu.com")
	stdout, err := cmd.CombinedOutput()
	cmd.Start()
	if err == nil {
		fmt.Println(string(stdout))
	}

	// r := bufio.NewReader(stdout)
	// line, _, err := r.ReadLine()
	// if err != nil {
	// 	fmt.Println(line)
	// }

}
