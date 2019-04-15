package main

import "fmt"

func Myfunc(args ...int) { //
	fmt.Println(len(args))
	for _, v := range args {
		fmt.Println(v)
	}
}

func Myfunc01(a int, args ...int) { //不定参数放在后面，固定参数a一定要传参
	fmt.Println(a)
	for _, v := range args {
		fmt.Println(v)
	}

}
func main() {
	Myfunc(1)
	Myfunc(1, 4)
	Myfunc01(3, 3)
}
