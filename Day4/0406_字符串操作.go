package main

import(
	"fmt"
	"strings"
)

func main(){
	//contains："hello"中是否包含"he"
	fmt.Println(strings.Contains("hello","he"))
	
	//Join 组合
	str := []string{"hello","he","liu"}
	buf := strings.Join(str,"")
	fmt.Println(buf) //字符串
	
	//index, 查看位置,不存在返回-1
	fmt.Println(strings.Index("helloworld","world")) //字符串
	
	//repeat :重复
	buf = strings.Repeat("go",3)
	fmt.Println(buf) //字符串
	
	//split 以指定分割符分割
	buf = "hello world my name is go"
	str = strings.Split(buf, " ")
	fmt.Println(str[1])//切片
	
	//trim :去掉两头字符
	buf = strings.Trim("    hello word !         "," ")
	fmt.Printf("#%s#\n",buf) //字符串
	
	//fields去掉空格
	s := strings.Fields("    hello word !         ")
	fmt.Println(s) //切片
	
}