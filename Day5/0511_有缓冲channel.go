package main

import (
	"fmt"
	"time"
)

func main(){
	
	//创建有缓冲channel
	c := make(chan int,3)
	fmt.Println(len(c),cap(c))  //len:现有数据长度，cap:容量，无缓冲都为0
	
	go func(){
		for i := 0; i < 10; i ++{
			fmt.Println("main i = ",i)
			c <-i 
			fmt.Println(len(c),cap(c))
		}
	}()
	time.Sleep(2*time.Second)
	for i := 0; i < 10; i ++{
			num := <- c //通道放满就会阻塞
			fmt.Println("num = ",num)
	}
}