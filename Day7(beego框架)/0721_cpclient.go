package main

import (
	"io/ioutil"
	"net"
	"os"
	"fmt"
)

func main(){
	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    //程序将用户的输入作为参数service传入net.ResolveTCPAddr获取一个tcpAddr,
    tcpAddr, err := net.ResolveTCPAddr("tcp4",service) 
    checkErr(err)
    
    //然后把tcpAddr传入DialTCP后创建了一个TCP连接conn
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkErr(err)
    
    //通过conn来发送请求信息
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkErr(err)
    
    //最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息。
    result, err := ioutil.ReadAll(conn)
    checkErr(err)

    
    fmt.Println(string(result))
    os.Exit(0)
   
}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}