package main

import (
	"github.com/astaxie/beego"
	"io"
	"net/http"
	"os"
	"log"
)
type myhandle struct{}

func (*myhandle)ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "URL:" + r.URL.String())
}
func main(){
	mux := http.NewServeMux()
	wd, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	mux.Handle("/",&myhandle{})
	mux.HandleFunc("/hello", sayhello)
	mux.Handle("/static/",
	http.StripPrefix("/staic/",http.FileServer(http.Dir(wd))))
	http.ListenAndServe(":8080", mux)
}

func sayhello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "version2")
	beego.
}
