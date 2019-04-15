package main

import "fmt"

type person struct{
	name string
	sex  string
	age  int
}

type student struct{
	person //匿名字段，继承了person所有的字段
	id   int
	addr string
}
func main(){
	
	//顺序初始化
	var s student = student{person{"he","liu",2},12,"a"}
	fmt.Println(s)
	
	//自动初始化
	s1 := student{person{"he","liu",2},23,"a"}
	//%+v 显示更详细的信息
	fmt.Printf("%+v",s1)
	
	//部分初始化
	s2 := student{person:person{name:"he"},id:1}
	fmt.Println(s2)
	
}