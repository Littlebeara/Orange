//In line ABCD, which ones of them have syntax issues?
package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B 错 
	f(p) //C
	//g(p) //D 不能在参数为接口指针的地方传入结构体的指针
}
