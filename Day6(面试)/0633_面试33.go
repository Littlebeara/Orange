package main

import "fmt"

func main() {
	list := new([]int) //不能编译通过，new不能用在这里，这里要用make
	list = append(list, 1)
	fmt.Println(list)
}