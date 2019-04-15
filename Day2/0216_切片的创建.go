package main

import "fmt"

func main() {
	//自动推导
	s := []int{1, 2, 3}
	fmt.Println(len(s), cap(s))

	//使用make函数
	c := make([]int, 5, 10) //类型，长度，容量（可以省略）
	fmt.Println(len(c), cap(c))

	c1 := make([]int, 5) //类型，长度，容量
	fmt.Println(len(c1), cap(c1))

}
