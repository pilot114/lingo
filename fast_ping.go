package main

import (
	// "github.com/tatsushid/go-fastping"
	// "os"
	"fmt"
	// "net"
	// "time"
)


func main() {
	aOctet := 123;
	for  bOctet:= 0; bOctet < 256; bOctet++ {
		for cOctet := 0; cOctet < 256; cOctet++ {
			for dOctet := 0; dOctet < 256; dOctet++ {
				fmt.Println(fmt.Sprintf("%d.%d.%d.%d", aOctet, bOctet, cOctet, dOctet))
			}
		}
	}

	// p := fastping.NewPinger()
	// ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// p.AddIPAddr(ra)
	// p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
	// 	fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	// }
	// p.OnIdle = func() {
	// 	fmt.Println("finish")
	// }
	// err = p.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }	
}

