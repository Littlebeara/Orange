package main

import (
	"fmt"
	"time"
)

func main(){
	
	//无缓冲channel
	c := make(chan int)
	
	go func(){
		for i := 0; i < 3; i ++{
			fmt.Println("i = ", i)
			num := <- c
			fmt.Println("num = ",num)
		}
	}()
	time.Sleep(2*time.Second)
	for i := 0; i < 3; i ++{
		fmt.Println("main i = ",i)
		c <-i //没有数据会阻塞
	}
}