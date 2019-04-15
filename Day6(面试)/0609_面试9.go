package main

import (
	"sync"
	"fmt"
	
)
var wg sync.WaitGroup
const N = 10

func test(m map[int]int,i int){
	mu := &sync.Mutex{} //Mutex提供了互斥锁的机制
	defer wg.Done() ////任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
	mu.Lock()
	m[i] = i
	mu.Unlock()
}
func main() {
	m := make(map[int]int)
	//WaitGroup可以等待一系列goroutine的执行完成。在main goroutine中通过
	//Add方法指定等待的个数，在每个goroutine中调用Done方法标记该goroutine执行
	//完成。Wait方法会等待，直到所有的goroutine执行结束。
	
	wg.Add(N)
	for i := 0; i <  N; i++ {
		test(m,i)
	}
	wg.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	fmt.Println(len(m),m)
}