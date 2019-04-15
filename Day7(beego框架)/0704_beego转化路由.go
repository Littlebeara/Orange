package main

import (
	"reflect"
	"net/url"
	"strings"
	"runtime"
	"net/http"
)

func( p *ControllerRegister) ServerHTTP (w http.ResponseWriter, r http.Request){
	defer func(){
		if err := recover(); err != nil{
			if !RecoverPanic{
				panic(err)
			}else{
				Critical("handler crashed with erroe", err)
				for i := 1; ;i +=1{
					_, file, line, ok := runtime.Caller(i)
					if !ok{
						break
					}
					Critical(file,line)
				}
			}
		}
	}()
	
	var started bool
	for prefix, staticDir := range StaticDir {
		if strings.HasPrefix(r.URL.Path, prefix){
			file := staticDir + r.URL.Path[len(prefix):]
			http.ServeFile(w,r,file)
			started = true
			return
		}
	}
	
	requestpath := r.URL.Path
	for _, route := range p.routers{
		
		if !route.regex.Matchstring(requestpath){
			continue
		}
		
		matches := route.regex.FindStringSubmatch(requestPath)
		if len(matches[0]) != len(requestpath){
			continue
		}
		
		params := make(map[string]string)
		if len (route.params) >0 {
			values := r.URL.Query()
			for i, match = range matches[1:]
			values.Add(route.params[1],match)
			params[route.params[i]]
		}
		
		r.URL.RawPath = url.Values(values).Encode() + "&" + r.URL.RawQuery
	}
	
	vc := reflect.New(route.controllerType)
	init := vc.MethodByName("init")
	in := make([]reflect.Value,2)
	ct := &Context{Responsewrier: w, request: r, Params: params}
	in[0] = reflect.ValueOf(ct)
	in[1] = reflect.ValueOf(route.controllerType.Name())
	init.Call(in)
	
	in = make([]reflect.Value, 0)
	method := vc.MethodByName("prepaer")
	method.Call(in)
	
	if r.Method == "GET" {
		method = vc.MethodByName("Get")
		method.Call(in)
	}else if r.Method == "POST" {
            method = vc.MethodByName("Post")
            method.Call(in)
        } else if r.Method == "HEAD" {
            method = vc.MethodByName("Head")
            method.Call(in)
        } else if r.Method == "DELETE" {
            method = vc.MethodByName("Delete")
            method.Call(in)
        } else if r.Method == "PUT" {
            method = vc.MethodByName("Put")
            method.Call(in)
        } else if r.Method == "PATCH" {
            method = vc.MethodByName("Patch")
            method.Call(in)
        } else if r.Method == "OPTIONS" {
            method = vc.MethodByName("Options")
            method.Call(in)
        }
		
		if Autorender {
			method = vc.MethodByName("render")
			method.Call(in)
		}
	
		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	 if started == false {
        http.NotFound(w, r)
    }
	
	
	
}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	