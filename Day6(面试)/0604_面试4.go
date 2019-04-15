package main

type S struct {
	m string
}

func f() *S { //返回的是一个地址
	return nil //A
}

func main() {
	p := S{"foo"}  //B
	print(p.m) //print "foo"
}
