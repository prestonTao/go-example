package main

import "net"

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// remotAddr, err := net.ResolveUDPAddr("udp", "100.64.195.116:9981")
	remotAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9981")
	chk(err)
	locaAddr, err := net.ResolveUDPAddr("udp", ":9980")
	chk(err)
	conn, err := net.ListenUDP("udp", locaAddr)
	chk(err)
	_, err = conn.WriteToUDP([]byte("hello world"), remotAddr)
	chk(err)

	// conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	// chk(err)
	// _, err = conn.WriteToUDP([]byte("hello world"), &net.UDPAddr{IP: net.IP{127, 0, 0, 1}, Port: 9981})
	// chk(err)
}
