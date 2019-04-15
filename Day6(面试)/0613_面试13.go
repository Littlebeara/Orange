package main

import (
	"math/rand"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			locker.Lock()
			m[rand.Int()] = rand.Int()
			locker.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}