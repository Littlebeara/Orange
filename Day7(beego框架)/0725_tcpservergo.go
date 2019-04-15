package main

import (
	"strconv"
	"fmt"
	"os"
	"time"
	"net"
)

func main(){
	service := ":8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
	checkErr(err)
	
	listener, err := net.ListenTCP("tcp",tcpAddr)
	checkErr(err)
	
	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		go handleclient(conn)
	}
}
//通过把业务处理分离到函数handleClient，我们就可以进一步地实现多并发执行了
func handleclient(conn net.Conn){
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(2*time.Minute)) //设置超时
	request := make([]byte, 128) //set maxium request length to 128KB to prevent flood attack
	for {
		read_len, err := conn.Read(request)
		if err != nil{
			fmt.Println(err)
			break
		}
		if read_len == 0{
			break
		}else if string(request) == "timestamp"{
			daytime := strconv.FormatInt(time.Now().Unix(),10)
			conn.Write([]byte(daytime))
			
		}else{
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128)
		//清理request，因为conn.Read()会将新读取到的内容append到原内容之后。
	}
	
}

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
        
    }
}