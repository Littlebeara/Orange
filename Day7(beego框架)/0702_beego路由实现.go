package main

import (
	"github.com/astaxie/beego"
	"strings"
	"reflect"
	"regexp"
)

//们设计了两个数据类型controllerInfo(保存路径和对应的struct，这里是一个reflect.Type类型)
//和ControllerRegistor(routers是一个slice用来保存用户添加的路由信息，以及beego框架的应用信息)
type ControllerInfo struct{
	regex 			*regexp.Regexp
	params 			map[int]string
	controllerType  reflect.Type
}

type ControllerRegister struct{
	routers		[]*controllerInfo
	Application *App
}

//ControllerRegistor对外的接口函数有
func (p *ControllerRegister)Add(pattern string, c ControllerInfo){
	parts := strings.Split(pattern, "/")
	
	j := 0
	params :=make(map[int]string)
	for i, part := range parts{
		if strings.HasPrefix(part, ":"){
			expr := "([^/]+)"
			///a user may choose to override the defult expression
            // similar to expressjs: ‘/user/:id([0-9]+)’
            if index := strings.Index(part,"("); index = -1{
            	expr = part[index:]
				part = part[:index]
				
            }
            params[j] = part
            part[i] = expr
            j ++ 
		}
	}
	//recreate the url pattern, with parameters replaced
    //by regular expressions. then compile the regex
    
    pattern = strings.Join(parts,"/")
    regex, regexErr := regexp.Compile(pattern)
    if regexErr != nil{
    	////TODO add error handling here to avoid panic
    	panic(regexErr)
		return
    }
    //now create the Route
    t := reflect.Indirect(reflect.ValueOf(c)).Type()
    route := &ControllerInfo{}
    route.regex = regex
    route.params = params
    route.controllerType = t
    
   p.routers = append(p.routers, route)
}

//上面我们实现的动态路由的实现，Go的http包默认支持静态文件处理FileServer，由于我们实现了自定义的路由器，
//那么静态文件也需要自己设定，beego的静态文件夹路径保存在全局变量StaticDir中，StaticDir是一个map类型.

func (app *App)SetStaticPath(url string, path string)*App{
	StaticDir[url] = path
	return app
}

beego.SetStaticPath("/img", "/static/img")
















