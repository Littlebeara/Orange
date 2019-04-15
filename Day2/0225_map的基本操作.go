package main

import "fmt"

func main() {
	s := map[int]string{1: "he", 2: "liu"}
	fmt.Println(s)
	s[1] = "liu" //可以修改
	fmt.Println(s)
	s[3] = "er" //可以添加
	fmt.Println(s)

	for _, v := range s {
		fmt.Println(v) //遍历map
	}
	//判断key是否存在
	//返回对应的值和key存在的状态 true or false

	value, ok := s[6]
	if ok == true {
		fmt.Println("存在", value)
	} else {
		fmt.Println("no")
	}

	delete(s, 1) //删除
	fmt.Println(s)

}
