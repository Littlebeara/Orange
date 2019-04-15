package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//1.生成种子,以时间为种子，因为时间会不断变化，如果种子是一样的，随机数也是一样的。
	rand.Seed(time.Now().UnixNano()) //获取当前纳秒

	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))

	}

}
