package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			println("通道输入", i)
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for nums := range ch {
			println("通道输出", nums)
		}
		println("通道已关闭")
	}()
	wg.Wait()
}
