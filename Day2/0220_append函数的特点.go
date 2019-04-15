package main

import "fmt"

func main() {
	a := make([]int, 1)
	c := cap(a)

	for i := 0; i < 8; i++ { //如果超过原来的容量以两倍增长
		a = append(a, i)
		if b := cap(a); b > c {
			fmt.Printf("%d-------------%d\n", c, b)
			c = b
		}
	}
}
