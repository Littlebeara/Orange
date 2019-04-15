package main

import(
	"fmt"
	"time"
)

func main(){
	timer := time.NewTimer(7 * time.Second)
	ok := timer.Reset(1 * time.Second)
	fmt.Println(ok)
	go func() {
		<- timer.C
		fmt.Println("结束")
	}()
	//timer.Stop() //停止
	for {
		
	}
}