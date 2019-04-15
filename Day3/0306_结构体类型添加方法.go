package main

import "fmt"

type student struct{
	name string
	id   int

}

func (p student)print(){
	fmt.Println(p)
}

func (p *student)set(name string,id int){
	p.name = "liu"
	p.id = 78
	
}
func main(){
	s := student{"he",23}
	s.print()
	
	(&s).set("he",78)
	s.print()
	
	
	
	
}