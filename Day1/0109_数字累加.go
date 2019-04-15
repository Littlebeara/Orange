package main

import "fmt"

func add(a int) int {
	if a == 100 {
		return 100

	}
	return a + add(a+1)
}
func main() {
	c := add(1)
	fmt.Println(c)

}
