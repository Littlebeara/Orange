package main

import "fmt"

func main() {
	a := 10
	str := "heliu"
	func() { //闭包以引用方式捕获外部变量
		a = 89
		str = "haha"
		fmt.Println(a, str)
	}()
	fmt.Println(a, str)

}
