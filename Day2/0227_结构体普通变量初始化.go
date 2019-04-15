package main

import "fmt"

func main() {
	type student struct {
		id   int
		name string
		sex  string
		age  int
	}

	//顺序初始化，每个成员都必须要初始化

	var s student = student{12, "he", "woman", 12}
	fmt.Println(s)

	//指定成员初始化
	s1 := student{id: 1, sex: "woman"}
	fmt.Println(s1)
}
