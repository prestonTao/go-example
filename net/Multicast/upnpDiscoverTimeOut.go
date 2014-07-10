package main

import (
	"fmt"
	"net"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	remotAddr, err := net.ResolveUDPAddr("udp", "239.255.255.250:1900")
	chk(err)
	locaAddr, err := net.ResolveUDPAddr("udp", "192.168.1.100:")
	chk(err)
	conn, err := net.ListenUDP("udp", locaAddr)
	chk(err)

	// msg := "GET /igd.xml HTTP/1.1\r\n" +
	// 	"User-Agent: Java/1.7.0_45\r\n" +
	// 	"Host: 192.168.1.1:1900\r\n" +
	// 	"Accept: text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2\r\n" +
	// 	"Connection: keep-alive\r\n\r\n"

	searchMessage := "M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"ST: urn:schemas-upnp-org:device:InternetGatewayDevice:1\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 3\r\n" + // seconds to delay response
		"\r\n"

	_, err = conn.WriteToUDP([]byte(searchMessage), remotAddr)
	chk(err)
	fmt.Println("send ok")
	for {
		buf := make([]byte, 1024)
		_, remotA, err := conn.ReadFromUDP(buf)
		chk(err)
		fmt.Println(string(buf), "\n  |  ", remotA.IP, "  |  ", remotA.Port)
	}

}
