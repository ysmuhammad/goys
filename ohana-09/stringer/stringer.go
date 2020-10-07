package main

import "fmt"

type IPAddr [4]byte
type Ys [4]string

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
func (y Ys) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", y[0], y[1], y[2], y[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	dv := Ys{"1", "2", "3", "4"}
	//dv := [4]string{"1","2","3","4"}
	fmt.Printf("%v", dv)
}
