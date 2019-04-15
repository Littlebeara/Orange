package main

import(
	"fmt"
	"encoding/json"
)

//首字母必须大写，生成json时
type IT struct{
	Company string   `json:"-"`         //不会输出到屏幕 :中间不要有空格
	Sujects []string `json:"subjects"` //二次编码
	IsOk    bool     `json:",string"`  //改成string类型
	Ptice   float64
}

func main(){
	it := IT{"jucuyun",[]string{"Python","Go","C"},true,23.2}
	//j, err := json.Marshal(it)
	j,err := json.MarshalIndent(it,""," ") //格式化
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(j)) //返回的是byte类型。要转换为字符串
}