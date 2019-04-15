package main

import(
	"fmt"
	"encoding/json"
)

type IT struct{
	Company string   `json:"company"`         //不会输出到屏幕 :中间不要有空格
	Price   float64	 `json:"price"`
	Sujects []string `json:"subjects"` //二次编码
	
}

func main(){
	var temp IT
	data := []byte( `{
    "company":"jucanyun",
    "price":34,
    "subjects":[
        "python",
        "java",
        "go"
    ]
}`)
	err := json.Unmarshal(data,&temp)//不要忘了&
	if err != nil{
		fmt.Println(err)
	}
	
	fmt.Println(temp)
	
	type IT struct{
		
	Company string   `json:"company"`         //不会输出到屏幕 :中间不要有空格
	
}
	
	var temp2 IT
	err = json.Unmarshal(data,&temp2)//不要忘了&
	if err != nil{
		fmt.Println(err)
	}
	
	fmt.Println(temp2)
	
	
	
}