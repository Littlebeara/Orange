package main

import "fmt"

func main() {
	a := map[int]string{}
	fmt.Println(a, len(a)) //没有长度，只有容量

	//可以通过make创建
	b := make(map[int]string)
	fmt.Println(b, len(b))

	//可以通过make创建,但是长度不是等于分配好的长度，而是里面实际的数据长度
	c := make(map[int]string, 10)
	c[1] = "heliu" //成员赋值 键是唯一的，不要重复
	fmt.Println(c, len(c))

}
