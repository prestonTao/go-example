// GWK_Example project main.go
package main

import (
	"fmt"
	"sdming-wk/wk"
)

func main() {
	// ./demo/basic/main.go for more detail

	server, err := wk.NewDefaultServer()
	
	if err != nil {
	    fmt.Println("NewDefaultServer error", err)
	    return
	}
	
	server.RouteTable.Get("/data/top/{count}").To(...)
	
	server.Start()
}
