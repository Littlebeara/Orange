package main

import (
	"fmt"
	"os"
	"html/template"
)

func main(){
	s1, _ := template.ParseFiles("header.tmpl","content.tmpl","footer.tmpl")
	s1.ExecuteTemplate(os.Stdout,"header",nil)
	fmt.Println()
	
	s1.ExecuteTemplate(os.Stdout,"content",nil)
	fmt.Println()
	
	s1.ExecuteTemplate(os.Stdout,"footer",nil)
	fmt.Println()
}