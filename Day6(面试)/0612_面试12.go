package main

import "fmt"
func main() {
	s := "123"
	ps := &s
	b := []byte(*ps) //把ps变成字节切片
	//fmt.Println(b)
	pb := &b

	s += "4"  //"1234"
	*ps += "5" //"12345"
	b[1] = '0' 
	fmt.Println(b)
	println(*ps) //"12345"
	println(string(*pb)) 
}