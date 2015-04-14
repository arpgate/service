package arpgate

import (
	"fmt"
	"github.com/ccding/go-stun/stun"
)

func CheckStun() int {

	// stun.SetServerHost("stun.arpgate.com", 3478)
	stun.SetServerHost("s1.taraba.net", 3478)
	//stun.SetServerHost("stun1.l.google.com", 19302)

	//stun.SetServerHost("stun1.voiceeclipse.net", 3478)

	nat, host, err := stun.Discover()
	if err != nil {
		fmt.Println(err)
	}

	switch nat {
	case stun.NAT_ERROR:
		fmt.Println("Test failed")
	case stun.NAT_UNKNOWN:
		fmt.Println("Unexpected response from the STUN server")
	case stun.NAT_BLOCKED:
		fmt.Println("UDP is blocked")
	case stun.NAT_FULL:
		fmt.Println("Full cone NAT")
	case stun.NAT_SYMETRIC:
		fmt.Println("Symetric NAT")
	case stun.NAT_RESTRICTED:
		fmt.Println("Restricted NAT")
	case stun.NAT_PORT_RESTRICTED:
		fmt.Println("Port restricted NAT")
	case stun.NAT_NONE:
		fmt.Println("Not behind a NAT")
	case stun.NAT_SYMETRIC_UDP_FIREWALL:
		fmt.Println("Symetric UDP firewall")
	}

	if host != nil {
		fmt.Println(host.Family())
		fmt.Println(host.Ip())
		fmt.Println(host.Port())
	}
	return nat
}
