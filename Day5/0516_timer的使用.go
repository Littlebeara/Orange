package main

import(
	"fmt"
	"time"
)
func main(){
	//创建一个定时器，设置时间为两秒，这个只会响应一次
	timer := time.NewTimer(2*time.Second)
	fmt.Println(time.Now())
	
	t := <- timer.C
	fmt.Println(t)
}