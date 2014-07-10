package main

import (
	"bytes"
	"fmt"
)

func main() {
	demo()
}

// <?xml version="1.0"?>
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
//   <SOAP-ENV:Body>
//     <m:GetStatusInfo xmlns:m="urn:schemas-upnp-org:service:WANIPConnection:1">
//     </m:GetStatusInfo>
//   </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>

func demo() {
	root := Node{Name: "SOAP-ENV:Envelope", Attr: map[string]string{"xmlns:SOAP-ENV": "@Service", "name": "tao"}}
	body := Node{Name: "SOAP-ENV:Body"}
	root.AddChild(body)
	xml := root.BuildXML()
	fmt.Println(xml)
}

// type NodeInterface interface {
// 	addChild(NodeInterface)
// 	buildXML() bytes.Buffer
// }

type Node struct {
	Name  string
	Attr  map[string]string
	Child []Node
}

func (n *Node) AddChild(node Node) {
	n.Child = append(n.Child, node)
}
func (n *Node) BuildXML() string {
	buf := bytes.NewBufferString("<")
	buf.WriteString(n.Name)
	for key, value := range n.Attr {
		buf.WriteString(" ")
		buf.WriteString(key + "=" + value)
	}
	buf.WriteString(">")
	for _, node := range n.Child {
		buf.WriteString(node.BuildXML())
	}
	buf.WriteString("</" + n.Name + ">")
	return buf.String()
}
