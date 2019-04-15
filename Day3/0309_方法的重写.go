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

//student也有一个同名的方法，这个叫做方法的重写
func (s *student)print(){
	fmt.Println(s)
}
func main(){
	var s1 student = student{person{"he",23},12,"bj"} 
	//就近原则，调用本作用域的方法，找不到再找继承过来的方法。
	s1.print()
	
	//显式调用
	s1.person.print()
}