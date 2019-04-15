package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for _, v := range a {
		fmt.Println(v)
	}

}
