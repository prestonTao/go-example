package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

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
