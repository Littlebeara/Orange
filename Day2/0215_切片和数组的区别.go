package main

import "fmt"

func main() {
	a := [5]int{} //中括号里的值是固定的，长度和容量也是固定的
	fmt.Println(len(a), cap(a))

	b := []int{} //中括号里的指为空或者是... 长度和容量不固定
	fmt.Println(len(b), cap(b))

	s := append(b, 11) //在切片末尾追加一个成员
	fmt.Println(len(s), cap(s))

}
