package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 76, 7, 8, 98}
	fmt.Println(len(array), cap(array))

	a := array[:] //从下标为0到len-1的数值
	fmt.Println(len(a), cap(a))

	b := array[:7] //从下标为0到7-1的数值
	fmt.Println(len(b), cap(b))

	c := array[4:] //从下标为4到len-1的数值
	fmt.Println(len(c), cap(c))

	d := array[2:5:8] //从下标为2到len-1的数值
	fmt.Println(len(d), cap(d))

}
