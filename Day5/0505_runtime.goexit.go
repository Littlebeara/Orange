package main

import (
	"fmt"
	"runtime"
)

func test(){
	fmt.Println("aaa")
	runtime.Goexit() //终止所在的协程
	fmt.Println("bbbb")
	
}

func main(){
	go func(){
		fmt.Println("ccc")
		test()
		fmt.Println("ddd")
	}()
	
	for { //无限循环，主协程不退出
		
	}
}