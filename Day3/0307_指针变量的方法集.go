package main

import "fmt"

type student struct{
	name string
	id   int
	
}

func(p *student)setvalue(){
	fmt.Println("setvalue")
}

func(p student)setpointer(){
	fmt.Println("setpointer")
}
func main(){
	s := &student{"he",12}
	s.setpointer()
	s.setvalue() //指针也可以传递给非指针的参数,会在内部做转换
	
	s1 := student{"liu",23}
	s1.setpointer() //非指针也可以传递给指针的参数，内部转换
	s1.setvalue()
	
}