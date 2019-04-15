package main

import (
	"time"
	"fmt"
)

func main(){
	ticker := time.NewTicker(2*time.Second) //可以循环，像闹钟
	
	i := 0
	for {
		<- ticker.C
		i ++
		fmt.Println(i)
		if i == 5{
			ticker.Stop()
			break
		}
	}
}