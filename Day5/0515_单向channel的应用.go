package main

import(
	"fmt"
)

func product(p1 chan<- int){
	for i := 0; i < 10; i ++{
		p1 <- i*i
	}
	close(p1)
}

func customer(p <- chan int){
	for data := range p {
		fmt.Println(data)
	}
}
func main(){
	ch := make(chan int)
	go product(ch)
	customer(ch)
}