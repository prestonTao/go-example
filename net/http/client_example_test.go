package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

/*
	自定义请求head，body，method，参数用body传递
	获取添加金币记录
*/
func TestGetAddCoin(t *testing.T) {
	url := "/findcoin"
	method := "POST"
	params := map[string]interface{}{
		"Page": 0,
		"Size": 10,
	}

	header := http.Header{"RANGE": []string{"0-200"},
		"User-Agent": []string{"OperatingPlatform"},
		"Accept":     []string{"text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2"}}
	client := &http.Client{}
	//req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	bs, err := json.Marshal(params)
	req, err := http.NewRequest(method, "http://127.0.0.1:8080"+url, strings.NewReader(string(bs)))
	if err != nil {
		fmt.Println("创建request错误")
		return
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求服务器错误")
		return
	}
	fmt.Println("response:", resp.StatusCode)
	if resp.StatusCode == 200 {
		robots, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取body内容错误")
			return
		}
		fmt.Println(string(robots))
	}
}

/*
	模拟网页form表单提交请求
*/
func TestPostRequest(t *testing.T) {
	data := make(url.Values)
	data["address"] = []string{"127.0.0.1:9981"}

	resp, _ := http.PostForm("http://127.0.0.1:9981/add", data)
	if resp.StatusCode == 200 {
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("返回结果：", string(body))
}

/*
	普通http请求，没有任何参数
*/
func TestOrdinaryClient(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:10010/university/1/major?access_token=haha")
	if resp.StatusCode == 200 {
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("返回结果：", string(body))
}

/*
	查看网关是否支持upnp协议
*/
func TestSsdp(t *testing.T) {
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

	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println(err)
		return
	}
	//bodystr := string(body)
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.Header)
		fmt.Println(string(body))
	}
}

/*
	请求自定义header
*/
func TestHeader(t *testing.T) {
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
