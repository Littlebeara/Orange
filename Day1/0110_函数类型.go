package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型

type functype func(int, int) int //没有{}

func main() {
	var a functype
	a = Add //后面不加括号
	c := a(1, 2)
	fmt.Println(c)

	var b functype
	b = Minus
	d := b(2, 1)
	fmt.Println(d)

}
