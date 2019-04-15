package main

import (
	"fmt"
	"time"
)

func main(){
	go func(){
		i := 0 
		for {
			i ++
			fmt.Println("test i =",i)
			time.Sleep(time.Second)
			}
		
	}() //主协程已经退出了，子协程还没来得及调用
}