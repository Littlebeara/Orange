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
	name string
}
func main(){
	var a student
	a.name = "he" //就近原则，操作的是student的name.如果student没有name字段，操作的就是person里面的字段
	a.person.name = "liu" //显示调用。操作的是person里面的字段
	a.id = 2
	fmt.Printf("%+v",a)
}