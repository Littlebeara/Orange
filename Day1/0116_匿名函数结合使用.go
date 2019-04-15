package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func() {
		fmt.Println(a, b) //函数结束前立即调用
	}()
	a = 0
	b = 2
	fmt.Println(a, b)

	c := 10
	d := 20
	defer func(c, d int) {
		fmt.Println(c, d) //事先已经把c,d的值放进来了。
	}(c, d)
	c = 0
	d = 32
	fmt.Println(c, d)
}
