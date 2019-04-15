package main

import "fmt"

func modify(a [5]int) { //数组做函数传递时值传递
	a[1] = 4
	fmt.Println(a)
}
func main() {
	a := [5]int{1, 3, 4, 65, 7}
	modify(a)
	fmt.Println(a)
}
