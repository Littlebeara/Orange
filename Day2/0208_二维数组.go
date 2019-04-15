package main

import "fmt"

func main() {
	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}

	}
	fmt.Println(a)

	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 65}}
	fmt.Println(b)

	c := [3][4]int{{2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(c)

	d := [3][4]int{1: {2, 3, 4}}
	fmt.Println(d)

}
