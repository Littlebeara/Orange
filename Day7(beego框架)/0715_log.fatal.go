package main

import (
	"fmt"
	"log"
)

func test_deferfatal(){
	defer func(){
		fmt.Println("---first---")
	}()
	log.Fatal("test for defer fatal") //对于 log.Fatal 接口，会先将日志内容打印到标准输
	//出，接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1 。但是有一点需要注意，由于是直
	//接调用系统接口退出，defer函数不会被调用，
}

func main(){
	test_deferfatal()
}