package main

import "fmt"

type student struct {
	id   int
	name string
	sex  int
}

func main() {
	var a student
	var p *student
	p = &a //不要忘了

	//通过指针操作成员，两种方式p.name 和(*p.id)p
	p.id = 1
	(*p).name = "he"
	p.sex = 3
	fmt.Println(*p)

	//通过new申请一个结构体 分配空间
	p1 := new(student)
	p1.id = 2
	(*p1).name = "he"
	p1.sex = 6
	fmt.Println(*p1)

}
