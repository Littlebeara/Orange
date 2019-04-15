package main 

import (
	"fmt"
	"time"
)

func main(){
	//time.After(2 * time.Second)
	//time.Sleep(2 * time.Second)
	timer := time.NewTimer(2 * time.Second)
	<- timer.C
	fmt.Println("时间到")
}