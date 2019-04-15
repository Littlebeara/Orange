package main

import (
	"fmt"
	"regexp"
)

func main(){
	buf := "43.3 rfgefc 4.5 7.0 89.09 ded"
	
	reg := regexp.MustCompile(`\d+.\d+`)
	
	res := reg.FindAllStringSubmatch(buf,-1)
	fmt.Println(res)
}