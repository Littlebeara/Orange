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

func shishi(i hunmaner){
	i.say()
}
func main(){
	//定义接口的类型
	//调用同一函数，不同表现，多态，多种形态
	s := &student{"he"}
	shishi(s)
	
	t := &teacher{3}
	shishi(t)
	
	//通过切片实现多态
	slice := make([]hunmaner,2)
	slice[0] = s
	slice[1] = t
	for _,v := range slice{
		v.say()
	}
	
	
	
	
	
	
	
	
	
	
	
	
	
}         