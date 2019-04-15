package main

import "fmt"

type student struct {
	id   int
	name string
	sex  string
}

func test(s student) {
	s.id = 22
	fmt.Println(s)
}

func main() {

	s := student{1, "he", "ge"}
	test(s)
	fmt.Println(s)
}
