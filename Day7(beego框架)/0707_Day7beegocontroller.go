package main

import (
	"path"
	"html/template"
)

type Controller struct{
	Ct			*Context
	Tpl			*template.Template
	Data		map[interface{}]interface{}
	Childname   string
	Tplnames	string
	Layout		[]string
	TplExt		string
	
}

type ControllerInterface interface{
	init(ct *context, cn string)
	Prepaer()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
	Render() error
}

Func (c *Controller) Init(ct *Context, cn string){
	c.Data = make(map[interface{}]interface{})
	c.Layout = make(string[], 0)
	c.Tplname = ""
	c.Childname = ch
	c.Ct = ct
	c.TplExt = "tpl"
}

func (c *Controller) Prepaer(){
	
}

func (c *Controller)Finish(){
	
}
func (c *Controller) Get() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
    http.Error(c.Ct.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller)Render()error{
	if (len(c.Layout)) > 0 {
		var filenames []string
		for _, file := range c.Layout{
			filenames = append(filesnames, path.Join(Viewspath, file))
		}
		t, err := template.ParseFiles(filenames...)
		if err != nil{
			Trace("template ParseFiles err:", err)
		}
		err = t.ExecuteTemplate(c.Ct.ResponseWriter, c.TplNames, c.Data)
		if err != nil{
			Trace("template Execute err:", err)
		}
	}else{
		if c.Tplnames = ""{
			c.Tplnames = c.Childname + "/" + c.Ct.Request.Method + "." + c.TplExt
		}
		 t, err := template.ParseFiles(path.Join(ViewsPath, c.TplNames))
        if err != nil {
            Trace("template ParseFiles err:", err)
        }
        err = t.Execute(c.Ct.ResponseWriter, c.Data)
        if err != nil {
            Trace("template Execute err:", err)
        }
    }
    return nil
}

func (c *Controller) Redirect(url string, code int) {
    c.Ct.Redirect(code, url)
}   
}























