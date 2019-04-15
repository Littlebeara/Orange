package main

import (
	"fmt"
	"regexp"
	
)

func main(){
	buf := "abc a9c a2c qtc xta"
	
	//1,解析规则
	
	//reg := regexp.MustCompile(`a.c`)
	reg := regexp.MustCompile(`a\dc`)
	
	res := reg.FindAllStringSubmatch(buf,-1)
	
	fmt.Println(res)
	
}