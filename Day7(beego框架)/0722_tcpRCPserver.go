package main

import (
	"os"
	"fmt"
	"net"
	"net/rpc"
	"errors"
)

type Args struct {
	A,B int
}

type Quotient struct{
	Quo, rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int)error{
	*reply = args.A * args.B
	return nil
	
}

func (t *Arith)Divide(args *Args, quo *Quotient)error{
	if args.B == 0{
		return errors.New("divide by zero")
	}
	
	quo.Quo = args.A / args.B 
	quo.rem = args.A % args.B
	return nil
}
func main(){
	arith := new(Arith)
	rpc.Register(arith)
	
	tcpaddr, err := net.ResolveTCPAddr("tcp",":1234")
	checkErr(err)
	
	listener, err := net.ListenTCP("tcp",tcpaddr)
	checkErr(err)
	
	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		rpc.ServeConn(conn)
		
	}
}

func checkErr(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
        
    }
}























