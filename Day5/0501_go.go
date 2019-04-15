package main

import (
	"fmt"
	"time"
)

func test (){
	for {
		fmt.Println("this is a new task")
		time.Sleep(time.Second)
	}
}

func main(){
	
	go test()
	for {
		fmt.Println("this is a main task")
		time.Sleep(time.Second)
	}
	
}