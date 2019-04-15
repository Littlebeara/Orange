package main

import (
	"fmt"
	"os"
)

func main(){
	
	//标准设备默认已经打开，用户可以直接。关闭后，无法写内容
	//os.Stdout.Close()  
	fmt.Println("hello") //往标准设备（屏幕）写内容
	
	
	os.Stdin.Close()  //关闭后无法输入
	var a int
	fmt.Scan(&a) //输入内容
	fmt.Println(a)
	
	
	
}