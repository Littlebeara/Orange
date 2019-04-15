package main

import "fmt"

type functype func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}
func calc(a, b int, functype functype) (res int) {

	res = functype(a, b)
	return

}

func main() {
	res := calc(2, 3, Add)
	fmt.Println(res)

	res1 := calc(2, 3, Minus)
	fmt.Println(res1)
}
