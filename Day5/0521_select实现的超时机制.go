package main

import(
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	quit := make(chan bool)
	go func (){
	for {
		select{
			case num := <- ch :
			fmt.Println(num)
			case <- time.After(3 * time.Second): //三秒之后停止
			
			quit <- true 
			
		}
	}
	
	}()
	<- quit
	fmt.Println("结束")
}