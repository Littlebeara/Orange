package main

type S struct{} //空结构体

func (s S) F() {} //方法为空  不懂

type IF interface {
	F() //接口定义了方法
}

func InitType() S { //返回的是结构体
	var s S //s 结构体  为空
	return s
}

func InitPointer() *S { //返回的是结构体指针 true
	var s *S
	return s
}
func InitEfaceType() interface{} { //返回的是空接口 false
	var s S
	return s
}

func InitEfacePointer() interface{} { //false
	var s *S
	return s
}

func InitIfaceType() IF { //接口 false
	var s S //s 结构体
	return s
}

func InitIfacePointer() IF { //返回值是接口 false
	var s *S //s 结构体指针
	return s
}

func main() {
	//println(InitType() == nil) cannot convet nil to type S
	println(InitPointer() == nil)  // true
	println(InitPointer())
	
	println(InitEfaceType() == nil)  //false
	println(InitEfaceType()) 
	
	println(InitEfacePointer() == nil) //false
	println(InitEfacePointer())
	
	println(InitIfaceType() == nil) //false
	println(InitIfaceType())
	
	println(InitIfacePointer() == nil) //false
	println(InitIfacePointer())
}

                        