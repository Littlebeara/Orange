package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	j := `{
    "company":"jucanyun",
    "price":34,
    "subjects":[
        "python",
        "java",
        "go"
    ]
}`
	m := make(map[string]interface{},4)
	err := json.Unmarshal([]byte(j),&m)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(m)
	
	//要用类型断言才能得到对应的value类型和值
	for i , v := range m {
		switch v.(type){
			case string:
			fmt.Println(i,v)
			case int:
			fmt.Println(i,v)
			case []interface{}: //这里定义的事空接口的切片类型
			fmt.Println(i,v)
		}
	}
	
	
	
	
	
	
}