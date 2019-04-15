package main

import "fmt"

type person struct{
	name string
	age  int
	sex  string
}

type student struct{
	*person
	id int
}
func main(){
	s := student{&person{"he",12,"df"},3}//不要忘了&号
	fmt.Println(s.age,s.id)
	
	//先定义变量
	var a student
	a.person = new(person) //后分配空间
	a.name = "liu"
	fmt.Println(a.name)
}