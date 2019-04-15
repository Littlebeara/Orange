package main

import "fmt"

func main(){
	c := make(chan string)
	defer fmt.Println("主协程调用完毕")
	go func(){
		defer fmt.Println("子协程结束")
		for i := 0 ; i < 5; i ++ {
		fmt.Println("i = ", i)
		}
		c <- "-------------"
		
	}()
	str := <- c
	fmt.Println(str)
}