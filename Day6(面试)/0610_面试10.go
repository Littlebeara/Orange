package main
 
import (
	"fmt"
	//"sync"
	"time"
)
 
//var waitgroup sync.WaitGroup
const N = 10
func main() {
	//waitgroup.Add(N)
	//mu := &sync.Mutex{}
	for i := 0; i <  N; i++ {
		go func() {
			//defer waitgroup.Done() ////任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
			//mu.Lock()
			a := i
			fmt.Println(i,&a)
			//mu.Unlock()
		}()
	}
	//waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	fmt.Println("done!")
	time.Sleep(2*time.Second)
	
}