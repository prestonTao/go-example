package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	packageSize int = 1024 //byte
)

func main() {
	connBaidu()
}

func connBaidu() {
	//resp, err := http.Get("http://www.baidu.com/")
	resp, err := http.Get("http://127.0.0.1/view?id=backup.tar.gz")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
			fmt.Println(robots)
		}
		//html := string(robots)
		for key, v := range resp.Header {
			fmt.Println("key:", key, "------------value:", v)
		}
	}
}
