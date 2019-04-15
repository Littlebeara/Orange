package main

import (
	"log"
	"html"
	"fmt"
	"net/http"
)

func foolhandler(w http.ResponseWriter, r http.Request){
	fmt.Fprintf(w, "hello, %q",html.EscapeString(r.URL.Path))
	
}

http.HandleFunc("/foo", foolhandler())

func handfunc("/bar", func(w ,http.ResponseWriter, r http.Request){
	fmt.Fprintf(w,"hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080",nil))