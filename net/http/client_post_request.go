package main

import (
	"net"
)

func main() {
	request()
}

/*
	模拟网页form表单提交请求
*/
func request() {
	data := make(url.Values)
	data["address"] = []string{"127.0.0.1:9981"}

	resp, _ := http.PostForm("http://127.0.0.1:9981/add", data)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("返回结果：", string(body))
}

/*
	普通http请求，没有任何参数
*/
func getList() {
	resp, _ := http.Get("http://127.0.0.1:10010/university/1/major?access_token=haha")
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("返回结果：", string(body))
}
