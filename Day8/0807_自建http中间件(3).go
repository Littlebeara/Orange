package main

import (
	"net/http"
)

type Appendhandle struct{
	hanler http.Handler
}

func (this *Appendhandle) ServeHTTP(w http.ResponseWriter, r *http.Request){
	this.hanler.ServeHTTP(w, r)
	w.Write([]byte("hello world3"))
}

func myhandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world4"))
}
func main(){
	mid := &Appendhandle{http.HandlerFunc(myhandler))
	http.ListenAndServe(":8081", mid)
}