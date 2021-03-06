package main

import (
	"fmt"
	"html/template"
)

func main(){
	tOk := template.New("first")
	template.Must(tOk.Parse("some static text /* and a comment */ "))
	//模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写。
	fmt.Println("The first one parsed OK")
	
	template.Must(template.New("second").Parse("some static text {{.Name}}"))
	fmt.Println("The second one parsed OK")
	
	fmt.Println("The next one ought to fail.")
	terr := template.New("err")
	template.Must(terr.Parse(" some statis text {{.name}"))
}