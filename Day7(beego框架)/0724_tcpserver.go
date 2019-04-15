package main

import (
	"time"
	"net"
	"fmt"
	"os"
)

func main(){
	service := ":7777"
	tcpaddr, err := net.ResolveTCPAddr("tcp4",service)
	checkErr(err)
	
	listener, err := net.ListenTCP("tcp",tcpaddr)
	checkErr(err)
	
	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}


