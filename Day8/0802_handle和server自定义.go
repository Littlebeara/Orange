package main

import (
	"log"
	"time"
	"io"
	"net/http"
	"os"
)
type myhandle struct{}

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func (*myhandle)ServeHTTP(w http.ResponseWriter, r *http.Request){
	if h, ok := mux[r.URL.String()]; ok{
		h(w,r)
		return
	}
	io.WriteString(w, "URL:" + r.URL.String())
}

func sayhello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello, version3")
}

func saybye(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "bye, version3")
}


func main(){

	server := http.Server{
		Addr 		: ":8080",
		Handler 	: &myhandle{},
		ReadTimeout : 5*time.Second,
	}
	
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayhello
	mux["/bye"]   = saybye
	
	
	err = server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
	
}

