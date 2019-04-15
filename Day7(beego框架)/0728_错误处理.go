package main

import (
	"fmt"
	"errors"
)


func sqrt(f float64)(float64, error){
	
	if f < 0{
		return 0, errors.New("math: square root of negative number")
	}
}
func main(){
	f, err := sqrt(-1)
	if err != nil{
		fmt.Println(err)
	}
}