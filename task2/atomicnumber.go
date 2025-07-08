package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var couter int64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&couter, 1)
				println("协程%s", id, "开始执行", couter)
			}
		}(i)
	}
	wg.Wait()

}
