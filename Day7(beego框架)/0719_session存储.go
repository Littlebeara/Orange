package main

import (
	"container/list" //list是一个双向链表。该结构具有链表的所有功能。
	"github.com/astaxie/session"
	"sync"
	"time"
)

var pder = &provider{list:list.New()} //创建一个新的list

type SessionStore struct{
	sid 		 string  					//session id 唯一标识
	timeAccessed time.Time					//最后访问时间
	value        map[interface]interface    //session里面存储的值
}

func(st *SessionStore)Set(key, value interface{})error{
	st.value[key] = value
	pder.sessionupdate(st.sid)
	return nil
}

func (st *SessionStore)Get(key interface{})interface{}{
	pder.Sessionupdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}else{
		return nil
	}
	return nil
}

func (st *SessionStore)Delete(key interface{})error{
	delete(st.value,key)
	pder.Sessionupdate(st.sid)
	return nil
}

func (st *SessionStore)SessionID()string{
	return st.sid
}

type provider struct{
	lock		sync.Mutex
	sessions 	map[string]*list.Element //在元素中存储的值,用来存储在内存
	list		*list.List				 //用来做gc
	
}

func ( pder *provider) SessionInit(sid string)(session.Session, error){
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{},0)
	newsess := &SessionStore{sid: sid, timeAccessed:time.Now(),value:v}
	element := pder.list.PushBack(newsess) //在list l的末尾插入值为v的元素，并返回该元素
	pder.sessions[sid] = element
	return newsess, nil
	
}

func(pder *provider) SessionRead(sid string)(session.Sessions, error){
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore),nil
	}else{
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *provider) SessionDestroy(sid string)error{
	if element, ok := pder.sessions[sid]; ok{
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *provider) SessionGC(maxlifetime int64){
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back() //获取list l的最后一个元素
		if element == nil{
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime < time.Now().Unix(){
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		}else{
			break
		}
	}
	
}

func (pder *provider) Sessionupdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element) //将元素e移动到list l的首部，如果e不属于list l，则list不改变。
		return nil
	}
	return nil
}

func init(){
	pder.sessions = make(map[string]*list.Element,0)
	session.register("memeory",pder)
}

































