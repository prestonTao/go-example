package main

import (
	// "fmt"
	// "io"
	"log"
	"net"
	"strings"
	// "os"
)

func main() {
	remotAddr, err := net.ResolveUDPAddr("udp", "224.0.1.251:5353")
	if err != nil {
		log.Println("组播：组播地址格式不正确")
	}

	// locaAddr, err := net.ResolveUDPAddr("udp", "192.168.1.107:9999")
	// if err != nil {
	// 	log.Println("组播：本地ip地址格式不正确")
	// }
	// conn, err := net.ListenUDP("udp", locaAddr)
	// conn, err := net.ListenUDP("udp", &net.UDPAddr{
	// 	IP:   net.IPv4zero,
	// 	Port: 11990,
	// })
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(getLocalIntenetIp()),
		Port: 0,
	})
	defer conn.Close()
	if err != nil {
		log.Println("组播：监听udp出错")
	}
	_, err = conn.WriteToUDP([]byte("hhahahahah"), remotAddr)
	if err != nil {
		log.Println("组播：发送msg到组播地址出错")
	}

}

//获取本机能联网的ip地址
func getLocalIntenetIp() string {
	/*
	  获得所有本机地址
	  判断能联网的ip地址
	*/

	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip
}
