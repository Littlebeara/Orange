//What will be printed when the code below is executed?
//What will be the exit code after the code below is executed?
//Copy

package main

import "log"

func f() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover:%#v", r) //打印了panic(1)的值
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}