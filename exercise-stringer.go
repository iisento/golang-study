package main

import "fmt"

// IPAddr ...
// 4桁のアドレスを格納
type IPAddr [4]byte

// Stinger ...
type Stinger interface {
	String()
}

// "."で区切った形式への変換
func (ip IPAddr) String() string {
	return fmt.Sprint(ip[0], ".", ip[1], ".", ip[2], ".", ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
