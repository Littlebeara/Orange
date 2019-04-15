package main

import "fmt"

type student struct {
	id   int
	name string
	sex  string
}

func main() {
	s := student{1, "he", "g"}
	s1 := student{1, "he", "g"}
	s2 := student{2, "he", "g"}

	fmt.Println(s == s1, s1 == s2)

	var s3 student
	s3 = s
	fmt.Println(s3)
}
