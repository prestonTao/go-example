package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	httpServerStartUP()
}

func startTCPClient(proto, host, url string, header http.Header) string {
	client := &http.Client{}
	// 第三个参数设置body部分
	// reqest, _ := http.NewRequest("GET", url, nil)

	// reqest.Proto = proto
	// reqest.Host = "http://" + host
	// reqest.Header = header

	// ----------------------
	reqest, _ := http.NewRequest("GET", "http://127.0.0.1"+url, nil)
	reqest.Proto = proto
	reqest.Host = "127.0.0.1"

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	// reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")
	// reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")
	reqest.Header.Set("Host", "127.0.0.1")
	// reqest.Header.Set("Referer", "http://baidu.com/")
	reqest.Header.Set("User-Agent", "tao")

	// ----------------------

	response, _ := client.Do(reqest)

	resultBody, _ := ioutil.ReadAll(response.Body)
	//bodystr := string(body)
	log.Println("http请求返回", response.StatusCode)
	if response.StatusCode == 200 {
		// log.Println(response.Header)
		return string(resultBody)
		// log.Println(string(resultBody))
	}
	return "error tao"
}

func httpServerStartUP() {

	mux := &MaizeMux{}

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
	fmt.Println("webServer startup...")
}

func hello(w http.ResponseWriter, r *http.Request) {

}

type MaizeMux struct{}

func (m *MaizeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := r.URL.Path
	// scheme := r.URL.Scheme
	proto := r.Proto
	host := r.Host
	header := r.Header
	port := r.RequestURI
	url := r.URL

	log.Println(path, "|", proto, "|", host, "|", header, "|", port, "|", url)

	io.WriteString(w, startTCPClient(proto, host, path, header))

}
