package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	hello()
	unmarshalToStruct()
}

//----------------------------------------------------------------------
//       第二个例子
//----------------------------------------------------------------------
//将xml文件解析为一个struct
func unmarshalToStruct() {
	content, err := ioutil.ReadFile("studygolang.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}

type Result struct {
	Person []Person
}
type Person struct {
	Name      string
	Age       int
	Career    string
	Interests Interests
}
type Interests struct {
	Interest []string
}

//----------------------------------------------------------------------
//       第一个例子
//----------------------------------------------------------------------
//解析xml文件
func hello() {
	var t xml.Token
	var err error

	input := `<Person type="main"><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`
	input = `<Person type="main"><FirstName>Xu</FirstName><LastName>Xinhua</LastName><test testtype="tt"/></Person>`
	inputReader := strings.NewReader(input)

	// 从文件读取，如可以如下：
	// content, err := ioutil.ReadFile("studygolang.xml")
	// decoder := xml.NewDecoder(bytes.NewBuffer(content))

	decoder := xml.NewDecoder(inputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			fmt.Println("开始标记: ", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Println("标记中的元素:", attrName, attrValue)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Println("结束标记：", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Println("标记中的文本：", content)
		default:
			// ...
		}
	}
}
