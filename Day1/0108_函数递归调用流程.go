package main

import "fmt"

func test(a int) {

	if a == 1 {
		return
	}
	test(a - 1)
	fmt.Println(a)

}

func main() {
	test(3)
}
