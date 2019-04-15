package main

import "fmt"

//面向对象， 
type long int

func (tmp long)test01(e long)long{
	return tmp - e
}

//面向过程 
func test(a, b int)int{
	return a +  b
}
func main(){
	a,b := 1,2
	c := test(a,b)
	
	var f long = 5
	//调用方法：变量名.函数名(参数)
	d := f.test01(5)
	fmt.Println(c, d)
}