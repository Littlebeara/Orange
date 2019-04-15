package main

import "fmt"

type student struct {
	id   int
	name string
	sex  string
	age  int
}

//.操作
func main() {
	var s student
	s.id = 1
	s.age = 18
	s.name = "heliu"
	fmt.Println(s)
}
