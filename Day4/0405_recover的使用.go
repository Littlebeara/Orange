package main

import "fmt"

func test(){
	fmt.Println("aa")
}
func test1(x int){
	defer func(){ //recover要结合defer使用 recover可以打印panic信息
		if err := recover(); err != nil{
		fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 10 //如果x > 9 内部会调用panic函数
}
func test2(){
	fmt.Println("aa")
}
func main(){
	test()
	test1(12)
	test2()
}