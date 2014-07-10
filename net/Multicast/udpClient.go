package main

import "net"
import "fmt"

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// remotAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:1900")
	// chk(err)

	// addr, err := net.ResolveUDPAddr("udp", ":9981")
	// chk(err)
	// conn, err := net.ListenUDP("udp", addr)
	// chk(err)
	// defer conn.Close()
	// _, err = conn.WriteToUDP([]byte("M-SEARCH * HTTP/1.1\r\n"+
	// 	"HOST: 239.255.255.250:1900\r\n"+
	// 	"MAN: \"ssdp:discover\"\r\n"+
	// 	"MX: 3\r\n"+
	// 	"ST: ssdp:rootdevice\r\n\r\n:1"), remotAddr)
	// chk(err)
	// buf := make([]byte, 1024)
	// _, remoteAddr, err := conn.ReadFromUDP(buf)
	// chk(err)
	// fmt.Println(string(buf), "\n  |  ", remoteAddr.IP, "  |  ", remoteAddr.Port)

	remotAddr, err := net.ResolveUDPAddr("udp", "239.255.255.250:1900")
	chk(err)
	locaAddr, err := net.ResolveUDPAddr("udp", "192.168.1.100:9980")
	chk(err)
	conn, err := net.ListenUDP("udp", locaAddr)
	chk(err)

	_, err = conn.WriteToUDP([]byte("M-SEARCH * HTTP/1.1\r\n"+
		"HOST: 239.255.255.250:1900\r\n"+
		"ST: urn:schemas-upnp-org:device:InternetGatewayDevice:1\r\n"+
		"MAN: \"ssdp:discover\"\r\n"+
		"MX: 3\r\n"+
		"\r\n"), remotAddr)
	chk(err)
	buf := make([]byte, 1024)
	_, remoteAddr, err := conn.ReadFromUDP(buf)
	chk(err)
	fmt.Println(string(buf), "\n  |  ", remoteAddr.IP, "  |  ", remoteAddr.Port)
}
