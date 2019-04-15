package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
)

func readfile(p string){
	//1.打开文件
	f,err := os.Open(p)
	if err != nil{
		fmt.Println(err)
		return
	}
	//2. defer 关闭文件
	defer f.Close()
	
	//3.按行读取
	//先建缓冲区
	b := bufio.NewReader(f)
	
	for {
		a, err := b.ReadBytes('\n')
			if err != nil{
				if err == io.EOF{
					break
				}
			fmt.Println(err)
			}
			fmt.Println(string(a))
		
		

	}
}
func main(){
	path:= "./demo.txt"
	readfile(path)
}