package main

import "fmt"

func test() int {
	var a int
	a++
	return a * a //普通函数用完，变量就释放
}

func test01() func() int { //返回值是一个函数类型
	var a int
	return func() int {
		a++
		return a * a //不关注捕获的变量是否已经超过了变量域，只要闭包还在使用，这些变量就还会存在。
		//闭包中变量的生命周期不由变量的作用域决定
	}

}
func main() {

	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println("-------我是分割线---------")

	f := test01()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
