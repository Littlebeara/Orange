package main

import (
	"time"
	"net/http"
	"fmt"
	
)

http.SetCookie(w http.ResponseWriter, cookie http.Cookie)

type cookie struct{
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	Maxage     int
	Secure     bool
	Httponly   bool
	Raw		   string
	Unpared    []string
	
	
}



func main(){
	//如何设置cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1,0,0)
	cookie := http.Cookie{Name:"username", Value:"heliu", Expires: expiration}
	http.SetCookie(w, &cookie)
	
	//如何读取cookie,第一种方式
	cookie, _ := r.cookie("username")
	fmt.Fprint(w,cookie)
	
	//第二种方式
	for _, cookie = range r.cookies(){
		fmt.Fprint(w,cookie.name)
	}
	
	
	
	
	
	
	
	

}




