package main

import (
	"encoding/xml"
	"fmt"
)

func main() {

	result := new(ResultXML)
	// result.Title = "html"
	result.Value = "gfas45fda1s5f"
	content, _ := xml.Marshal(result)
	fmt.Println(string(content))
}

type ResultXML struct {
	XMLName xml.Name `xml:"html"`
	Value   string   `xml:"body>textarea"`
}
