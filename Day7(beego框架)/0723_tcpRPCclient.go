package main

import (
	"log"
	"net/rpc"
	"fmt"
	"os"
)



type Args struct{
	A, B int
}

type Quotient struct{
	Quo, Rem int
}

func main(){
	if len(os.Args) != 2{
		fmt.Println("Usage :", os.Args[0], " server:port")
		os.Exit(1)
	}
	sevice := os.Args[1]
	client, err := rpc.Dial("tcp", sevice)
	if err != nil{
		log.Fatal("dialing : ", err)
	}
	
	args := Args{17,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil{
		log.Fatal("Arith Multiply Error :", err)
	}
	fmt.Printf("Multiply: %d * %d = %d\n", args.A, args.B, reply)
	
	var Quot Quotient
	err = client.Call("Arith.Divide", a /
	rgs, &Quot)
	if err != nil{
		log.Fatal("Arith Divide Error :", err)
	}
	fmt.Printf("Divide:%d/%d=%d remainder %d\n", args.A, args.B, Quot.Quo, Quot.Rem)
	
	
}







