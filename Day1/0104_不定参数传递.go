package main

import "fmt"

func tmp(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
}

func tmp02(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
}
func test(args ...int) {
	tmp(args...)
}

func test01(args ...int) {
	tmp02(args[1:]...) //把后面两个参数传过去
}
func main() {
	test(0, 2, 3)
	test01(0, 2, 3)
}
