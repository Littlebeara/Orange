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
		switch v.(type){
			case int:
				fmt.Println(i,v)
			case string:
				fmt.Println(i,v)
			case student:
				fmt.Println(i,v)
		}
	}
	}