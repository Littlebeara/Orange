package main

import "fmt"

func mm(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max, min = b, a
	}
	return max, min
}
func main() {
	max, min := mm(10, 3)
	fmt.Println(max, min)
}
