package main

import "fmt"

func main() {
	//比较，只支持==和!=
	b := [5]int{1, 3, 5, 3, 3}
	c := [5]int{1, 3, 5, 3, 3}
	d := [5]int{1, 3, 5, 3}

	fmt.Println(b == c, c == d)

	//赋值
	var a [5]int
	a = b
	fmt.Println(a)
}
