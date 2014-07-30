package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	example1()
}

func example1() {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	fmt.Println(hostname)
	ip, err := net.LookupCNAME(hostname)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ip)
	addr, _ := net.LookupHost("localhost")
	fmt.Println(addr)

	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		ip, _, _ := net.ParseCIDR(a.String())
		fmt.Println(a, "============ip:", ip)
	}
}
