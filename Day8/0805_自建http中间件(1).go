package main

import (
	"net/http"
)

type SingleHost struct{
	handler    http.Handler
	allowedhost string
}

//中间键
func (this *SingleHost)ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Host == this.allowedhost{
		this.handler.ServeHTTP(w,r)
	}else{
		w.WriteHeader(403)
	}
}

//正常逻辑部分
func myhandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world!"))
}
func main(){
	singe := &SingleHost{
		handler		: http.HandlerFunc(myhandler),
		allowedhost : "localhost:8081",
	}
	http.ListenAndServe(":8081", singe)
}