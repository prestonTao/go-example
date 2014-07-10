package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"time"
)

var (
	packageSize int = 200 //byte
)

func main() {
	connBaidu()
	//conn()
}

func connBaidu() {
	resp, err := http.Get("http://www.baidu.com/")
	//resp, err := http.Get("http://127.0.0.1/view?id=backup.tar.gz")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		defer resp.Body.Close()
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
		fmt.Println(robots)
	}
}

func conn() {
	//header := http.Header{"Date": []string{time.Now().String()},
	//	"Accept":          []string{"image/gif, image/x-xbitmap, image/jpeg, image/pjpeg, application/vnd.ms-excel, application/msword, application/vnd.ms-powerpoint, */*"},
	//	"Accept-Language": []string{"zh-cn"},
	//	"Accept-Encoding": []string{"gzip, deflate"},
	//	"Connection":      []string{"Keep-Alive"}}

	header := http.Header{"RANGE": []string{"0-200"},
		"User-Agent": []string{"NetFox"},
		"Accept":     []string{"text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2"}}
	client := &http.Client{}
	//req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	req, err := http.NewRequest("GET", "http://127.0.0.1/view?id=backup.tar.gz", nil)
	if err != nil {
		fmt.Println("1111111111111")
	}
	req.Header = header
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("2222222222222")
	} else {
		defer resp.Body.Close()
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("3333333333333")
		}
		html := string(robots)
		fmt.Println(html)
	} else {

	}
}
