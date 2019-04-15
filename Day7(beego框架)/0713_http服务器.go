package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func  sayhello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()//解析参数，默认不会解析
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key :",k)
		fmt.Println("value : ",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello") //输出到客户端
	
	
}

func main(){
	http.HandleFunc("/",sayhello)//设置访问路由
	err := http.ListenAndServe(":9090",nil)
	if err!= nil{
		log.Fatal("listenandserver :", err)
	}
	
	
}
