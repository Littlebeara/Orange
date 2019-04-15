package main

import "fmt"
 
func xx(args ...interface{}){
	//可以接受任意类型和任意个数的参数
}
func main(){
	//空接口可以保持任何类型的值
	var i interface{} = 1
	fmt.Println(i)
	
	i = "asd"
	fmt.Println(i)

}