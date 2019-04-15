package main

import "fmt"
func main() {
	var m map[string]int //A：这里声明了变量，还需要初始化
	m = make(map[string]int)
	m["a"] = 1
	//if v := m["b"]; v != nil { //B： int类型不能专换为nil
	v , ok := m["b"]
	if ok == true{
		
		println(v)
	}
	fmt.Println(m)
}
