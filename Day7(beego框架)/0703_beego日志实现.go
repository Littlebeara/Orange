package main

import (
	"os"
	"log"
)

const{
	LevelTrace = iota
	levelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
}

var level = LevelTrace

func level()int{
	return level
}

func setlevel(l int){
	level = l
}

var BeeLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func SetLogger(l *log.Logger){
	BeeLogger = 1
}

func Trace (v ...interface{}){
	if level <= LevelTrace {
		BeeLogger.Printf("[T] %v\n", v)
	}
}

func Debug(v ...interface{}){
	if level <= LevelDebug{
		BeeLogger.Printf("[D] %v\n",v)
	}
}

func Info(v ...interface{}){
	if level <= LevelInfo{
		BeeLogger.Printf("[I] %v\n",v)
	}
}

func Warn(v ...interface{}){
	if level <= LevelWarn{
		BeeLogger.Printf("[W] %v\n",v)
	}
}

func Error(v ...interface{}){
	if level <= LevelError{
		BeeLogger.Printf("[E] %v\n",v)
	}
}


func Critical(v ...interface{}){
	if level <= LevelCritical{
		BeeLogger.Printf("[C] %v\n",v)
	}
}








































