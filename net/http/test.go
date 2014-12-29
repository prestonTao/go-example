package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	// "net/url"
	"strconv"
	"strings"
)

func main() {
	// http.Post("http://sso.easou.com/api2/validate.json:", bodyType, nil)

	// data := make(url.Values)
	// data["head"] = map[string]string{"sign": "085c740989cd3e47fe7c40cf9949af2f", "appId": "2073", "partnerId": "1000100010001028"}
	// data["body"] = map[string]string{"EASOUTGC": "8b1162276ac5f553e3a13fdf75b8ca6d"}

	// resp, _ := http.PostForm("http://sso.easou.com/api2/validate.json", data)
	// body, _ := ioutil.ReadAll(resp.Body)
	// log.Println(string(body), "=================================")

	readMappingBody := `{"head":{"sign":"085c740989cd3e47fe7c40cf9949af2f","appId":"2073","partnerId":"1000100010001028"},"body":{"EASOUTGC":"8b1162276ac5f553e3a13fdf75b8ca6d"}}`

	client := &http.Client{}
	// 第三个参数设置body部分
	reqest, _ := http.NewRequest("POST", "http://sso.easou.com/api2/validate.json", strings.NewReader(readMappingBody))
	reqest.Proto = "HTTP/1.1"
	reqest.Host = "http://sso.easou.com/api2/validate.json"

	reqest.Header.Set("Accept", "text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2")
	reqest.Header.Set("Content-Type", "text/xml")
	reqest.Header.Set("SOAPAction", "\"urn:schemas-upnp-org:service:WANIPConnection:1#GetStatusInfo\"")

	reqest.Header.Set("Connection", "Close")
	reqest.Header.Set("Content-Length", string(len([]byte(readMappingBody))))

	response, _ := client.Do(reqest)

	body, _ := ioutil.ReadAll(response.Body)
	//bodystr := string(body)
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println(response.Header)
		fmt.Println(string(body))
	}
}

//===================================================
var UPLOAD_DIR = "D:/test/"

func example1() {
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/file", File)
	http.ListenAndServe(":80", nil)
	// http.FileServer(root)
}

//文件上传
func Upload(w http.ResponseWriter, r *http.Request) {

}

func File(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	buf, _ := ioutil.ReadFile(UPLOAD_DIR + name)
	// w.Header().Set("Content-Type", "image")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	w.Header().Set("Content-Disposition", "attachment;filename="+name)
	w.Header().Set("Content-Type", "application")
	w.Write(buf)
	// http.ServeFile(w, r, imagePath)
}
