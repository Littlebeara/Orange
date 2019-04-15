package main

import (
	"strings"
	"fmt"
    "html/template"
    "os"
)

type Person1 struct {
    UserName string
}

type Friend struct{
	Fname string
}

type Person struct{
	UserName string
	Emails	 []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{})string{
	ok := false
	var s string
	if len(args) == 1{
		s, ok = args[0].(string)	
	}
	if !ok{
		s = fmt.Sprint(args...)
	}
	substr := strings.Split(s,"@")
	if len(substr) != 2{
		return s
	}
	return (substr[0] + " at " + substr[1])
	
	
}
func main() {
    t := template.New("fieldname example")
    t, _ = t.Parse("hello {{.UserName}}!") //{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象
    p := Person1{UserName: "Astaxie"}
    t.Execute(os.Stdout, p)
    
    
    f1 := Friend{Fname:"heliu"}
    f2 := Friend{Fname:"hejiaqi"}
    t  = template.New("exampple")
 	t  = t.Funcs(template.FuncMap{"emaildeal" : EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{.|emaildeal}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
            {{end}}
            `)
    s := Person{UserName:"he",
    	 Emails : []string{"heliu@gmail.com"," hejiaqi@gmail.com"},
		 Friends : []*Friend{&f1, &f2}}
	t.Execute(os.Stdout,s)
}