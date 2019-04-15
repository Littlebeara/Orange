package main

import (
	"net/http/httptest"
	"net/http"
)

type ModHandler struct{
	handler http.Handler
}

func (this *ModHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	rec := httptest.NewRecorder()
	this.handler.ServeHTTP(rec, r)
	
	for k,v := range rec.Header(){
		w.Header()[k] = v
	}
	
	w.Header().Set("go-web-foundation", "vip")
	w.WriteHeader(408)
	w.Write([]byte("test\n"))
	w.Write(rec.Body.Bytes())
	
}
func myhandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world5"))
}
func main(){
	mid := &ModHandler{http.HandlerFunc(myhandler)}
	http.ListenAndServe(":8081", mid)
}