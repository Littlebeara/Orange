package main

import (
	"fmt"
	"os"
)

func main(){
	os.Mkdir("heliu",0777)
	os.MkdirAll("heliu/test1/test2", 0777)
	err := os.Remove("heliu")
	if err != nil{
		fmt.Println(err)
	}
	os.RemoveAll("heliu")
	
	userfile := "heliu.txt"
	fout, err := os.Create(userfile)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i ++{
		fout.WriteString("just a test\r\n")
		fout.Write([]byte("just a test\r\n"))
	
	}
	
	f1, err := os.Open(userfile)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f1.Close()
	buf := make([]byte,1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n{
			break
		}
		os.Stdout.Write(buf[:n])
	}

	
}