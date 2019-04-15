package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args

	n := len(list)

	fmt.Println(n)

	for _, data := range list {
		fmt.Println(data)
	}

}
