package main

import "fmt"

func test(s map[int]string) { //map做函数参数时引用传递

	delete(s, 1)

}
func main() {
	s := map[int]string{1: "he", 2: "string", 3: "liu"}
	test(s)
	fmt.Println(s)
}
