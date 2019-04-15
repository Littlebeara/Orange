package main

import "fmt"

func tmp() int {
	return 77
}
func tmp1() (result int) {
	result = 55
	return
}
func main() {
	a := tmp()
	fmt.Println(a)
	result := tmp1()
	fmt.Println(result)
}
