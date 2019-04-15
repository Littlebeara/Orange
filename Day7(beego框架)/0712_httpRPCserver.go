package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"errors"
)


type Args struct{
	A,B int
}

type Quotient struct{
	Quo, Rem int
	
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
	
}

func (t *Arith)Divide(args *Args, quo *Quotient) error{
	if args.B == 0{
		return errors.New("divide by zero")
	}
	
	quo.Quo = args.A / args.B
	quo.Rem = args.A * args.B
	return nil
}
func main(){
	
	//我们注册了一个Arith的RPC服务，然后通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上，
	//然后我们就可以利用http的方式来传递数据了。
	airth := new(Arith)
	rpc.Register(airth)
	rpc.HandleHTTP()
	
	err := http.ListenAndServe(":1234",nil)
	if err != nil{
		fmt.Println(err.Error())
	}
	
}














