package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
	if (len(os.Args)) != 2{
		fmt.Fprintf(os.Stderr, "Usage : %s, ip-addr\n",os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	//ParseIP(s string) IP函数会把一个IPv4或者IPv6的地址转化成IP类型
	addr := net.ParseIP(name)
	if addr == nil{
		fmt.Println("invalid IP")
	}else{
		fmt.Println("the address is ", addr.String())
	}
	os.Exit(0)
}