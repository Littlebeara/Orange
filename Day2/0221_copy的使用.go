package main

import "fmt"

func main() {
	a := []int{12, 34, 45, 56}
	b := []int{234, 45, 456, 5, 6}
	copy(a, b)
	fmt.Println(a)
}
