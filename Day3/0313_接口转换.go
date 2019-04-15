package main

import "fmt"

type hunmaner interface{//子集
	sayhi()
}

type personer interface{ //超集
	hunmaner //会继承humaner的方法
	saysingg(s string)
}

type student struct{
	name string
	id   int
}

func (s *student)sayhi(){
	fmt.Printf("%s : hello \n",s.name)
}

func (s *student)saysingg(a string){
	fmt.Printf("%s sing : %s",s.name,a)
}
func main(){
	
	//超集可以转换为子集，子集不能转换为超集
	var i personer
	var h hunmaner
	i = &student{"he",23}
	//i = h //error
	h = i
	h.sayhi()
	
}