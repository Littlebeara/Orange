package main

import "fmt"

func main() {
	var a [50]int
	for i := 0; i < 50; i++ {
		a[i] = i + 1
		fmt.Println(a[i])
	}
	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
}
