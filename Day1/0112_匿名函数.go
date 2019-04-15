package main

import "fmt"

func main() {
	a := 10
	func() { //没有名字的函数
		fmt.Println(a)
	}()

	//有参数的匿名函数
	func(b int) {
		fmt.Println(b + b)
	}(20)

	//有参有返回值的匿名函数
	d := func(c int) int {
		return c - 6 //有返回值就要写return
	}(9)
	fmt.Println(d)

}
