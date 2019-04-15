package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// var peo People = Stduent{}
	s := Stduent{}
	think := "bitch"
	//fmt.Println(peo.Speak(think))
	fmt.Println(s.Speak(think)) //直接调用。golang的方法集仅仅影响接口实现和方法表达式转化，
	//与通过实例或者指针调用方法无关。
}