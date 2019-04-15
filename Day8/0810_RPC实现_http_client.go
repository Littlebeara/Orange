package main

import (
	"log"
	"net/rpc"
	"fmt"
	"os"
)

type Args struct {
	A,B int
}

type Quotient struct{
	Quo, Rem int
}

func main(){
	if len(os.Args) != 2{
		fmt.Println("Usage：", os.Args[0], "server")
		os.Exit(1)
	}
	serveraddr := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serveraddr+":8081")
	if err != nil{
		log.Fatal("dialog:", err)
	}
	args := Args{17,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil{
		log.Fatal("Call:", err)
	}
	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)
	var quo Quotient
	
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil{
		log.Fatal("call:", err)
	}
	fmt.Printf("%d / %d = %d\n", args.A, args.B, quo.Quo)
	fmt.Printf("rem:%d", quo.Rem)
}













