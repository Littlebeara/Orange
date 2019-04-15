package main

import(
	"fmt"
	"os"
	"io"
)

func writestring(p string){
	//1. 创建文件
	f, err := os.Create(p)
	if err != nil{
		fmt.Println(err)
	}
	
	//2. defer 关闭文件
	defer f.Close()
	
	//3.往文件写入内容
	for i := 0; i < 10; i ++{
		a := fmt.Sprintf("i = %d\n",i)
		n , err := f.WriteString(a)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(n)
	}
}

func readfile(p string){
	//1. 打开文件
	f, err := os.Open(p)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	//2.defer 关闭文件
	defer f.Close()
	
	//3.读文件
	var b = make([]byte,1024)
	n, err := f.Read(b)
	if err != nil  && err != io.EOF{
		fmt.Print(err)
		return
	}
	fmt.Println(string(b[:n]))
	
}
func main(){
	
	path := "./demo.txt"
	//writestring(path)
	readfile(path)
}