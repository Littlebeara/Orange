package main

import "fmt"

func main() {
	array := []int{}
	fmt.Println(array, len(array), cap(array))

	array = append(array, 1)
	array = append(array, 1)
	array = append(array, 1)
	fmt.Println(array, len(array), cap(array))

}
