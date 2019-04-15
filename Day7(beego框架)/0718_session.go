package main

import (
	"html/template"
	"time"
	"net/url"
	"net/http"
	"encoding/base64"
	"crypto/rand"
	"io"
	"fmt"
	"sync"
)


//定义一个全局的session管理器
type Manager struct{
	cookieName string
	lock       sync.Mutex
	provider   Provider
	maxlifetime int64
	
}

func Newmanager (ProvideName, cookieName string, maxlifetime int64)(*Manager, error){
	provider, ok := provides[ProvideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q(forgotten import?)",ProvideName )
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime : maxlifetime},nil
}


//Go实现整个的流程应该也是这样的，在main包中创建一个全局的session管理器


var globleSessions *session.Manger
//然后在init函数中初始化
func init(){
	globleSessions, _ = Newmanager("memory","gosessionid",3600)
}

//我们知道session是保存在服务器端的数据，它可以以任何的方式存储，比如存储在内存、数据库或者文件中。
//因此我们抽象出一个Provider接口，用以表征session管理器底层存储结构。
type provider interface{
	
	//SessionInit函数实现Session的初始化，操作成功则返回此新的Session变量
	SessionInit(sid string)(session, error)
	
	//SessionRead函数返回sid所代表的Session变量，如果不存在，
	//那么将以sid为参数调用SessionInit函数创建并返回一个新的Session变量
	SessionRead(sid string)(session, error)
	
	//SessionDestroy函数用来销毁sid对应的Session变量
	SessionDestory(sid string)error
	
	//SessionGC根据maxLifeTime来删除过期的数据
	SessionGC(maxlifetime int64)
	
}

//对Session的处理基本就 设置值、读取值、删除值以及获取当前sessionID这四个操作，
//所以我们的Session接口也就实现这四个操作。

type session interface{
	Set(key, value interface{})error
	Get(key, interface{}) interface{}
	Delete(key, interface{}) error
	sessionID()string
	
}

//注册存储session的结构的Register函数的实现
var provides = make (map[string]provider)

func Register(name string, provide provider){
	if provide == nil {
		panic("session:Resgister provide is nil")
	}
	if _, dup := provides[name]; dup{
		panic("session : Resgiter called twice for provide " + name)
	}
	
	provides[name] = provider
}

//Session ID是用来识别访问Web应用的每一个用户，因此必须保证它是全局唯一的（GUID)

func ( manager *Manager) SessionID() string{
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//我们需要为每个来访用户分配或获取与他相关连的Session，以便后面根据Session信息来验证操作。
//SessionStart这个函数就是用来检测是否已经有某个Session与当前来访用户发生了关联，如果没有则创建之。

func (manager *Manager) sessionstart(w, http.ResponseWriter, r *http.Request)(session session){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err!= nil ||cookie.Value = ""{
		sid = manager.SessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie = http.Cookie{Name:manager.cookieName, Value: url.QueryEscape(sid), Path:"/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w,&cookie)
	}else{
		sid, _ := url.QueryUnescape(cookie.Value)
		session,_ := manager.provider.sessionread(sid)
	}
	return
	
}

func count(w http.ResponseWriter, r *http.Request){
	
	sess := globleSessions.Sessionstart(w,r)
	createtime := sess.get("createtime")
	if createtime == nil{
		sess.Set("createtime",time.Now().Unix())
		}else if(createtime.(int64) + 360 ) < (time.Now().Unix()){
			globleSessions.Sessiondestroy(w,r)
			sess = globleSessions.Sessionstart(w,r)
	
		}
		ct := sess.Get("countnum")
		if ct == nil{
			sess.Set("countnum",1)
		}else{
			sess.Set("countnum",(ct.(int)+1))
		}
		
		
		t, _= template.ParseFiles("count.gtpl")
		w.Header().Set("content- type","text/html")		
		t.excute(w,sess.get("countnum"))
	
	}
	//session重置
	func (manager *Manager) SessionDestroy (w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie(manager.cookieName)
		if err != nil || cookie.Value = ""{
			return
		}else{
			manager.lock.Lock()
			defer manager.lock.Unlock()
			manager.provider.SessionDestroy(cookie, Value)
			expiration := time.Now()
			cookie := http.Cookie{Name:manager.cookieName, Path:"/",HttpOnly: true, Expires: expiration, MaxAge:-1}
			http.Cookie(w,r)
		}
		
		//session销毁
		
		func init(){
			go globleSessions.GC()
		}
		func (manager *Manager)GC(){
			manager.lock.Lock()
			defer manager.lock.Unlock()
			manager.provider.SessionGC(manager.maxlifetime)
			time.Afterfunc(time.Duration(manager.maxlifetime),func()) {manager.GC()}		
		}
		
}






