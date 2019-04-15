package main

import "fmt"

func test(){
	fmt.Println("aa")
}
func test1(){
	panic("panic test")
}
func test2(){
	fmt.Println("aa")
}
func main(){
	test()
	test1()
	test2()
}