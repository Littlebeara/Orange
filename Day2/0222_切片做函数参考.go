package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initdata(a []int) { //切片做函数参数直接是引用传递
	n := len(a)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}
	//fmt.Println(a)
}

func bubble(a []int) {
	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

}
func main() {
	n := 10

	a := make([]int, n)

	initdata(a)
	fmt.Println(a)
	bubble(a)
	fmt.Println(a)

}
