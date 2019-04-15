package main

import "fmt"

type person struct{
	name string
	age  int
	sex  string
}

type student struct{
	person
	id   int
	addr string
}
func main(){
	var a student
	a.name = "he"
	a.id = 2
	fmt.Println(a)
	
	var b student
	b.person = person{"he",12,"d"}
	fmt.Println(b)
}