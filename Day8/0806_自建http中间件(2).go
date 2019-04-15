package main

import (
	"net/http"
)

//中间件

func SingleHost(handler http.Handler, allowedhost string)http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		if r.Host == allowedhost{
			handler.ServeHTTP(w,r)
		}else{
			w.WriteHeader(403)
		}
	}
	return http.HandlerFunc(fn)
}

//正常逻辑
func myhandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world2"))
}
func main(){
	single := SingleHost(http.HandlerFunc(myhandler), "localhost:8081")
	http.ListenAndServe(":8081", single)
}