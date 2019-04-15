package main

import (
	"net/http"
	//"os"
	"fmt"
	//"net"
	"net/rpc"
	"errors"
)

type Args struct {
	A,B int
}

type Quotient struct{
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int)error{
	*reply = args.A * args.B
	return nil
	
}

func(t *Arith)Divide(args *Args, quo *Quotient)error{
	if args.B == 0{
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main(){
	math := new(Arith)
	rpc.Register(math)
	rpc.HandleHTTP()
	
	err := http.ListenAndServe(":8081", nil)
	if err != nil{
		fmt.Println(err.Error())
	}
	
}














