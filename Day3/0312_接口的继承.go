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
	
	//定义接口变量
	var i personer
	s := &student{"he",12}
	i = s
	i.sayhi()
	i.saysingg("我爱你")
}