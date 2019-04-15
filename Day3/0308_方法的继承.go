package main

import "fmt"

type person struct{
	name string
	age  int
}


//继承了成员名也继承了方法
type student struct{
	person
	id   int
	addr string
}

func (p *person)print(){
	fmt.Println(p.name,p.age)
}
func main(){
	var s student = student{person{"he",23},12,"bj"} 
	s.print()
}