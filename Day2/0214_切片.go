package main

import "fmt"

func main() {
	a := [...]int{
	1, 4,
	5, 76, 7}
	b := a[0:3:5]
	fmt.Println(b, len(b), cap(b))
	c := a[1:4:5]
	fmt.Println(c, len(c), cap(c))

}
