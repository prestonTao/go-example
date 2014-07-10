package main

import (
	// "bytes"
	"fmt"
	// "io"
	"io/ioutil"
	"net/http"
	// "reflect"
	"strings"
)

func main() {
	//demo()
	fmt.Println("-----------------------------------------------------")
	ssdp()
	//demo()
}

func ssdp() {

	readMappingBody := `<?xml version="1.0"?>
	<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	<SOAP-ENV:Body>
	<m:GetPortMappingNumberOfEntries xmlns:m="urn:schemas-upnp-org:service:WANIPConnection:1">
	</m:GetPortMappingNumberOfEntries></SOAP-ENV:Body></SOAP-ENV:Envelope>`

	// reader := strings.NewReader(readMappingBody)
	// reader := bytes.NewReader([]byte(readMappingBody))
	// var rea io.Reader = *reader

	// limiReader := io.LimitedReader{reader.(io.Reader), int64(len(readMappingBody.getByte()))}

	// fmt.Println(reflect.TypeOf(*reader))
	// var rea io.Reader = strings.NewReader(readMappingBody)

	client := &http.Client{}
	// 第三个参数设置body部分
	reqest, _ := http.NewRequest("POST", "http://192.168.1.1:1900/ipc", strings.NewReader(readMappingBody))
	reqest.Proto = "HTTP/1.1"
	reqest.Host = "192.168.1.1:1900"

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

func demo() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}
