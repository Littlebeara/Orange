package main

import "fmt"

func test() (int, int, int) {
	return 1, 2, 3
}
func test1() (a int, b int, c int) {
	a, b, c = 4, 5, 6
	return a, b, c
}
func main() {
	a, b, c := test()
	fmt.Println(a, b, c)
	d, e, f := test1()
	fmt.Println(d, e, f)
}
