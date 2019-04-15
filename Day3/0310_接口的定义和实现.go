package main

import "fmt"

type hunmaner interface{
	
	//方法，只有声明没有实现，由别的类型（自定义类型）实现
	say()
}

type student struct{
	name string
}

type teacher struct{
	id  int
}

func (s *student)say(){
	fmt.Println(s.name)
}

func (t *teacher)say(){
	fmt.Println(t.id)
}
func main(){
	
	//定义接口的类型
	var i hunmaner
	
	//只要实现了此接口方法的类型，那么这个类型的变量（接受者类型）就可以给i赋值
	s := &student{"he"}
	i = s
	i.say()
	
	t := &teacher{3}
	i = t
	i.say()
}