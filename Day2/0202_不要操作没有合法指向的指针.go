package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)

	*p = 999 //error，没有指向

}
