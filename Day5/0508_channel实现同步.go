package main

import (
	"fmt"
	"time"
)

var a = make(chan int)

func printer(str string){
	for _,data := range str {
		fmt.Print(string(data))
		time.Sleep(time.Second)
	}
	
}

func person1(){
	printer("hello")
	a <- 666
}

func person2(){
	<- a //没有数据，阻塞
	printer("world")
}

func main(){
	go person1()
	go person2()
	
	for{
		
	}
}