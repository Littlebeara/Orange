package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}/*
很经典的题！ 这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

var in interface{}
另一种如题目：

type People interface {
    Show()
}
他们的底层结构如下：

type eface struct {      //空接口
    _type *_type         //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      //带有方法的接口
    tab  *itab           //存储type信息还有结构实现方法的集合
    data unsafe.Pointer
iface比eface 中间多了一层itab结构。 itab 存储_ty
pe信息和[]fun方法集，从上面的结构我们就可得出，因为dat
a指向了nil 并不代表interface 是nil， 所以返回值并不为空，
这里的fun(方法集)定义了接口的接收规则，在编译的过程中需要验证是否实现接口