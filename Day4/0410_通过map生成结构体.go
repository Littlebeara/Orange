package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	m := make(map[string]interface{},4)
	m["company"] = "jucanyun"
	m["subjects"] = []string{"python","java","go"}
	m["price"] = 34
	res, err := json.MarshalIndent(m,"","	")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(res))
}

