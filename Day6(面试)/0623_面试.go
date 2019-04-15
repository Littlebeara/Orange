//Fill in the blanks in line A and B, to parse terminal arguments to ip and port
//hich default to 0.0.0.0 and 8000.
package main

import "flag"
import "fmt"

var ip string
var port int

func init() {
	 flag.StringVar(&ip,"IP","0.0.0.0","default IP ") //a
	 flag.IntVar(&port,"port",8000,"default port")    //b
	
}

func main() {
	flag.Parse()
	fmt.Printf("%s:%d", ip, port)
}