package main

import "fmt"

func Myfunc01(a int) {
	fmt.Println(a)

}

func Myfunc02(a, b int) {
	fmt.Println(a, b)

}

func Myfunc03(a int, b string) {

	fmt.Println(a, b)

}

func main() {
	Myfunc01(666)
	Myfunc02(111, 222)
	Myfunc03(565, "heliu")
}
