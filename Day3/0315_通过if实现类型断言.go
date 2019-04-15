package main

import "fmt"

type student struct{
	name string
	id int
}

func main(){
	
	slice := make([]interface{},3)
	slice[1] =  2
	slice[2] = "sd"
	slice[0] = student{"he",23}
	for i, v := range slice{
		if value,ok:= v.(int); ok == true{
				fmt.Println(i,value)
		}else if value,ok:= v.(string); ok == true{    
			fmt.Println(i,value)
			}else  if value,ok:= v.(student);ok == true{    
			fmt.Println(i,value.name)
	}
}
}
