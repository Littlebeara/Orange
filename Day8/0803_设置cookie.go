package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/",cookie)
	http.HandleFunc("/2",cookie)
	http.ListenAndServe(":8081", nil)
}

func cookie(w http.ResponseWriter, r *http.Request){
	ck := &http.Cookie{
		Name  : "mycookie",
		Value : "hello",
		Path  :"/",
		Domain: "localhost",
		MaxAge: 120,
	}
	http.SetCookie(w, ck)
	
	ck2, err := r.Cookie("mycookie")
	if err != nil{
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}


func cookie2(w http.ResponseWriter, r *http.Request){
	ck2 := &http.Cookie{
		Name  : "mycookie",
		Value : "hello",
		Path  :"/",
		Domain: "localhost",
		MaxAge: 120,
	}
	
	w.Header().Set("set-cookie", ck2.String())
	ck2, err := r.Cookie("mycookie")
	if err != nil{
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}