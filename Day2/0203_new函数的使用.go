package main

import "fmt"

func main() {
	p := new(int) // 自动分配空间，自动推导变量
	*p = 66
	fmt.Println(p, *p)
}
