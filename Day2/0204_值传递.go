package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Println(a, b) //作用域不同，用完之后立马释放
}

func main() {
	a := 10
	b := 20

	swap(10, 20)
	fmt.Println(a, b)
}
