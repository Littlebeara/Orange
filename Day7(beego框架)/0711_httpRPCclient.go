package main

import (
	"log"
	"net/rpc"
	"fmt"
	"os"
)


type Args struct {
	A, B int
}

type Quotient struct{
	Quo, rem int
}

func main(){
	if len(os.Args) != 2{
		fmt.Println("Usage :", os.Args[0], "server")
		os.Exit(1)
	}
	
	serveraddress := os.Args[1]
	
	client, err:= rpc.DialHTTP("tcp", serveraddress + ":1234")
	if err != nil{
		log.Fatal("dialing", err)
	}
	
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	//在客户端作为client.Call的第2，3两个参数的类型。客户端最重要的就是这个Call函数
	//有3个参数，第1个要调用的函数的名字，第2个是要传递的参数，第3个要返回的参数(注意是指针类型)
	if err != nil{
		log.Fatal("Arith error:", err)
	}
	fmt.Printf("Arith :%d*%d = %d\n", args.A, args.B,reply)
	
	
	var quote Quotient
	err=client.Call("Arith.Divide",args, &quote)
	if err != nil{
		log.Fatal("Arith Divide", err)
	}
	fmt.Printf("Arith : %d/%d = %d remainder = %d\n", args.A, args.B,quote.Quo, quote.rem)
}




















