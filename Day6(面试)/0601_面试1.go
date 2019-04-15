package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10 //离开作用域后释放，没有改变值
	}

	for i := range ss {
		ss[i] += 10 //改变了值
	}

	//fmt.Println(ss)
	fmt.Println(s) //坑人
}