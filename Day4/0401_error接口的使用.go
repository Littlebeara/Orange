package main

import (
	"fmt"
	"errors"
)

func main(){
	err1 := fmt.Errorf("%s","this a normal error")
	err2 := errors.New("this a normal error2")
	fmt.Println(err1,err2)
}