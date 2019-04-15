package main

import(
	"fmt"
)
func feibo(a chan int, b chan bool){
	x,y := 1,1
	for {
		select{
			case a <- x:
				x, y = y, x+y
			case <- b :
				return
		}
	}
	
}
func main(){
	a := make(chan int)
	b := make(chan bool)
	go func() {
		for i := 0; i < 8; i ++ {
			fmt.Println( <- a)
		}
		b <- true
	}()
	feibo(a,b)
}