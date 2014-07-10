package main

import (
	"fmt"
	"log"
	"net"
)

// 组播的地址是保留的D类地址从224.0.0.0—239.255.255.255，
// 而且一些地址有特定的用处如，224.0.0.0—244.0.0.255只能用于局域网中路由器是不会转发的，
// 并且224.0.0.1是所有主机的地址，224.0.0.2所有路由器的地址，224.0.0.5所有ospf路由器的地址，
// 224.0.13事PIMv2路由器的地址；239.0.0.0—239.255.255.255是私有地址（如192.168.x..x）；
// 224.0.1.0—238.255.255.255可以用与Internet上的。
func main() {

	_net := "udp"
	ifis, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Couldn't get interface: %v -- Reason: %v", ifis, err)
	}
	for _, ifi := range ifis {
		log.Printf("FlagMulticast: ", ifi.Flags)
	}
	ifi := ifis[1]

	mcastAddrs, err := ifi.MulticastAddrs()
	if err != nil {
		log.Fatalf("Couldn't get mcastAddrs: %v -- Reason: %v", ifi, err)
	}
	log.Printf("MulticastAddrs: %s", mcastAddrs)
	// https://en.wikipedia.org/wiki/Multicast_address#Administratively_Scoped_IPv4_Multicast_addresses
	_mcastAddr := "239.0.0.1" // 239.0.0.0/8
	_mcastPort := 1234

	laddr := fmt.Sprintf("%s:%d", _mcastAddr, _mcastPort)

	gaddr, err := net.ResolveUDPAddr(_net, laddr)
	if err != nil {
		log.Fatalf("Couldn't resolve addr: %v (%v) -- Reason: %v", laddr, gaddr, err)
	}
	log.Printf("IsMulticast: %v", gaddr.IP.IsMulticast())
	con, err := net.ListenMulticastUDP(_net, nil, gaddr)
	if err != nil {
		log.Fatalf("Couldn't net.ListenMulticastUDP: %v -- Reason: %v", con, err)
	}
	log.Printf("Joined Multicast Group!")
	for {
		var buf = make([]byte, 1024)
		amt, remote, err := con.ReadFrom(buf)
		if err != nil {
			log.Fatalf("con.ReadFrom: %v", err)
		}
		log.Printf("Got: `%s` (%d) from %v", buf[:amt], amt, remote)
		response := fmt.Sprintf("Response! %s\n", buf[:amt])
		_, err = con.WriteTo([]byte(response), remote)
		if err != nil {
			log.Fatalf("con.WriteTo: %v", err)
		}
		log.Printf("Sent of: %#v to %s", response, remote)
	}

}
