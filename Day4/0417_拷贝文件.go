package main

import(
	"fmt"
	"io"
	"os"
	
)

func main(){
	//必须是三个
	list := os.Args
	if len(list) != 3{
		fmt.Println("请输入源文件和目的文件")
		return
	}
	
	srcfilename := list[1]
	detfilename := list[2]
	
	if srcfilename == detfilename{
		fmt.Println("不能相等")
		return
	}
	
	//源文件以只读方式打开
	sf,err := os.Open(srcfilename)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	
	//新建目的文件
	df,err := os.Create(detfilename)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	//关闭文件
	defer  df.Close()
	defer  sf.Close()
	
	
	//读取源文件内容，写入目的文件
	for { //记得for循环
		b := make([]byte,1024) //字符切片
		n, err := sf.Read(b)
		if err != nil{
			if err == io.EOF{
			break
		}
		fmt.Println(err)
	}
		df.Write(b[:n])
	}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
}