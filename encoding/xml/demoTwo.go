package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

func main() {
	unmarshalToStructAndTag()
}

//解析xml文件中的attribute
func unmarshalToStructAndTag() {
	content, err := ioutil.ReadFile("studygolangTwo.xml")
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

type Person struct {
	Name      string `xml:",attr"` //"attr"前面加逗号传说是支持小写，但我试了下大小写都能解析
	Age       int    `xml:",attr"`
	Career    string
	Interests Interests
}
type Result struct {
	Person []Person
}

type Interests struct {
	Interest []string
}

//这里可以自定义标记名称以及大小写

// type Person struct {
// 	Name      string    `xml:"name,attr"`
// 	Age       int       `xml:"age,attr"`
// 	Career    string    `xml:"career"`
// 	Interests Interests `xml:"interests"`
// }
// type Result struct {
// 	Person []Person `xml:"person"`
// }

// type Interests struct {
// 	Interest []string `xml:"interest"`
// }
