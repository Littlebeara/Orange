package main

import (
	"time"
	"fmt"
	"bytes"
	"io"
	"crypto/md5"
)

//登陆处理
func (this *Logincontroller) Post(){
	this.TPLnames = "login.tpl"
	this.Ctx.request.Parseform()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	 md5password := md5.New()
	io.WriteString(md5password,password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5password.Sum(nil))
	newpass := buffer.String()
	
	now := time.Now().Format("2006-01-02 15:04:05")
	
	UserInfo := models.GetUserInfo(username)
	if UserInfo.password == newpass{
		var users models.user
		users.last_logintime = now
		models.UpdateUserInfo(users)
		
		sess := globleSessions.Sessionstart(this.Ctx.ResponseWriter, this.Ctx.Request)
		sess.Set("uid", userInfo.Id)
        sess.Set("uname", userInfo.Username)

        this.Ctx.Redirect(302, "/")
	}
	
	////注册处理
	func (this *RegController)post(){
		this.Tplnames = "reg.tpl"
		this.Ctx.Request.ParseForm()
		username := this.Ctx.Request.Form.Get("username")
    	password := this.Ctx.Request.Form.Get("password")
		usererr = checkUsername(username)
		fmt.Println(usererr)
		if usererr == false{
			thid.Data["UsernameErr"] = "Username error, Please to again"
			return
		}
		
		passerr := checkPassword(password)
   		if passerr == false {
        this.Data["PasswordErr"] = "Password error, Please to again"
        return
   		}
    	md5password = md5.New()
		io.WriteString(md5password, password)
		buffer := bytes.NewBuffer(nil)
		fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
		
		now := time.Now().Format("2006-01-02 15:04:05")
		
		userInfo := models.GetuserInfo(username)
		
		if userInfo.Username = ""{
			var user model.user
			users.Username = username
       	 	users.Password = newPass
       		users.Created = now
       		users.Last_logintime = nowss
			models.Adduser(users)
			
		
		
			sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
        	sess.Set("uid", userInfo.Id)
       		sess.Set("uname", userInfo.Username)
       		this.Ctx.Redirect(302, "/")
   		} else {
		
			this.Data["UsernameErr"] = "User already exists"
		}
func checkPassword(password string) (b bool) {
    if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
        return false
    }
    return true
}

func checkUsername(username string) (b bool) {
    if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
        return false
    }
    return true
}
	
	
	//有了用户登陆和注册之后，其他模块的地方可以增加如下这样的用户是否登陆的判断：
	func (this *AddBlogController)prepare(){
		sess := globleSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		sess_uid := sess.Get("userid")
		sess_username := sess.Get("username")
		if sess_uid == nil{
			this.Ctx.Redirect(302, "/admin/login")
			return
		}
		this.Data["username"] = sess_username
	}
	
	
	
	
}
















