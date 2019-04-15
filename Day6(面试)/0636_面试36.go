package main

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return nil, false //nil 可以用作 interface、function、pointer、map、slice 和
	// channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型.
	//如果把nil改掉，他是会返回存在数据的
}
func main()  {
	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}

	v,err:=GetValue(intmap,3)
	fmt.Println(v,err)
}