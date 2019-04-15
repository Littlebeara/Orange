package main

import "fmt"

func main() {
	var a [5]int
	a = [5]int{1, 3, 5, 3, 3}
	fmt.Println(a)

	//自动推导类型，全部初始化
	b := [5]int{1, 3, 5, 3, 3}
	fmt.Println(b)

	//部分初始化
	c := [5]int{1, 3}
	fmt.Println(c)

	//指定初始化
	d := [5]int{1: 3, 3: 3}
	fmt.Println(d)

}
