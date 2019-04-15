package main

import (
	log "github.com/cihub/seelog"
	"log"
	"net/http"
)



func(p *Mymux)ServerHTTP(w http.ResponseWriter, r http.Request){
	if r.URL.Path == "/"{
		sayhello(w, r)
		return
	}
	NotFound404(w, r)
	return
}

func NotFound404(w http.ResponseWriter, r http.Request){
	log.Error("页面找不到") 
	t, _ := t.Parsefiles("tmpl/404.html", nil)
	Errorinfo := "文件找不到"
	t.Excute(w, Errorinfo)
	
}

func Systemerror(w http.ResponseWriter, r http.Request){
	log.Critical("系统错误")
	t, _ = t.ParseFiles("tmpEl/error.html", nil)
	Errorinfo := "系统暂时不可用"
	t.Excute(w, Errorinfo)
	
}