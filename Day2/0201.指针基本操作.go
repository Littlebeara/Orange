package main

import "fmt"

func main() {

	a := 10
	fmt.Println(a, &a)

	var p *int
	p = &a

	*p = 99
	fmt.Println(p, *p, a)
}
