package main

func main(){
	c := make(chan int) //默认双向
	
	var writechannel chan <- int = c  //双向可以转换为单向
	
	writechannel <- 999
	// <- 
	
	var readchannel <- chan int = c
	 <- readchannel
	//readchannel <- 88
	
	
	//var ch chan int =  writechannel //单向不能转换成双向
}