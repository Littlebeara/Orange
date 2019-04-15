package main

import "fmt"

func main() {
	defer fmt.Println("bbbb") //函数结束前调用
	fmt.Println("aaa")
}
