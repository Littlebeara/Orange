package main

import(
	"fmt"
	"runtime"
)

func main(){
	n := runtime.GOMAXPROCS(4) //以四核运算
	fmt.Println(n)
	
	for {
		go fmt.Print("1")
		fmt.Print("0")
	}
}
