package main

import (
	"strconv"
	"strings"
	"fmt"
)

func main(){
	//包含
	fmt.Println(strings.Contains("seafood","foo"))
	fmt.Println(strings.Contains("seafodd","we"))
	fmt.Println(strings.Contains("seafood",""))
	fmt.Println(strings.Contains("",""))
	
	//连接
	s := []string{"foo","d","sea"}
	fmt.Println(strings.Join(s,","))
	
	//位置
	fmt.Println(strings.Index("seafood","fo"))
	fmt.Println(strings.Index("seafood","mn"))
	
	//重复
	fmt.Println("se" + strings.Repeat("re",2))
	
	//替换
	
	fmt.Println("ok, ok, ok",strings.Replace("ok ok ok ","k","ok",2)) //替换两次
	fmt.Println("ok, ok, ok",strings.Replace("ok ok ok ","k","ok",-1)) //全部替换
	
	//分割
	fmt.Printf("%q\n", strings.Split("a,b,c",","))
	fmt.Printf("%q\n", strings.Split("a man a plan a banban","a"))
	fmt.Println("%q\n",strings.Split("xyc",""))
	fmt.Println("%q\n",strings.Split("","xyc"))
	
	//去掉指定字符串
	fmt.Printf("[%q]",strings.Trim(" !!!!! heliu !!  ! ","! "))
	
	//去掉空格并且以slice返回
	fmt.Printf("fields are %q", strings.Fields("  foo bui  ji"))
	
	//字符串转化的函数在strconv
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str,4567,10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '和') //注意中文的符号
	fmt.Println(string(str))
	
	//把数字转换为字符串
	e := strconv.Itoa(124)
	
	//字符串转换为数字
	d,err := strconv.Atoi("234")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(e,d)
	
	
	
	
	
	
	
	
	
	
	
	
	
	 
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	

	
	
	
	
	
	
	
	
	
}