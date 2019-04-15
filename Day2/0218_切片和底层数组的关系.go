package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5, 7, 788, 45, 56, 565}
	s := a[2:7]
	s[2] = 8 //切片的值改变，底层的值也会被改变
	fmt.Println(a)

	s1 := s[4:6:7]
	s1[1] = 8
	fmt.Println(a)

}
