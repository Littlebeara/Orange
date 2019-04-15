package main

import(
	"fmt"
	"errors"
)

func test(a,b int)(result int, err error){
	err = nil
	if b == 0 {
		err = errors.New("分母为0")
	}else{
		result = a/b
	}
	return
}

func main(){
	result, err := test(10,0)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}