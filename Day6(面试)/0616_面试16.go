package main

import "fmt"
type S struct {
	name string
}

func main() {
	m := map[string]S{"x": S{"one"}}
	//m["x"].name = "two"
	fmt.Println(m)//map[x:{one}]
	m["x"] = S{"two"}
	fmt.Println(m)
	
}

                        