package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createnum(p *int) {
	var num int
	rand.Seed(time.Now().UnixNano())
	for {
		num = rand.Intn(10000)
		if num > 1000 {
			break
		}
	}
	*p = num

}

func getnum(s []int, randnum int) {
	s[0] = randnum / 1000
	s[1] = randnum % 1000 / 100
	s[2] = randnum % 100 / 10
	s[3] = randnum % 10
}

func game(s []int) {
	var n int
	b := make([]int, 4)
	for {
		for {
			fmt.Println("请输入一个四位数")
			fmt.Scan(&n) //要取地址
			if n > 999 && n < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}

		getnum(b, n)
		d := 0
		for i := 0; i < 4; i++ {
			if s[i] > b[i] {
				fmt.Println("小")
			} else if s[i] < b[i] {
				fmt.Println("大")
			} else {
				fmt.Println("对")
				d++
			}

		}
		if d == 4 {
			fmt.Println("全对")
			break
		}
	}
}

func main() {
	var randnum int
	createnum(&randnum)
	//fmt.Println(randnum)
	s := make([]int, 4)
	getnum(s, randnum)
	//fmt.Println(s)
	game(s)

}
