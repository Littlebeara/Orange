package main

import (
	"os"
	"html/template"
)

func main(){
	tempty := template.New("test")
	tempty = template.Must(tempty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tempty.Execute(os.Stdout,nil)
	
	tWithValue := template.New("test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}}我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout,nil)
	
	tifelse := template.New("test")
	tifelse = template.Must(tifelse.Parse("if-else demo :{{if `anything`}} if 部分 {{else}} else部分 .{{end}}\n"))
	tifelse.Execute(os.Stdout,nil)
}