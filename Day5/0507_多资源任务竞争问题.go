package main

import (
	"fmt"
	"time"
)

func printer(str string){
	for _,data := range str {
		fmt.Print(string(data))
		time.Sleep(time.Second)
	}
	
}

func person1(){
	printer("hello")
}

func person2(){
	printer("world")
}

func main(){
	go person1()
	go person2()
	
	for{
		
	}
}