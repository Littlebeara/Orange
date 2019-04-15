package main

import "fmt"

const (
	x = iota //iota在const关键字出现时将被重置为0(const内部的第一行之前)，
	//const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}