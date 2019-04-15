package main

import (
	"fmt"
)

func main(){
	c := make(chan int,3)
	go func(){
		for i :=0; i < 3; i ++{
			c <- i 
		}
		close(c) //不需要写数据时，可以关闭，可以不用像关闭文件已经时时刻刻都要关闭
	}()
	
 for v := range c { //只有一个变量
	fmt.Println(v)
}
}