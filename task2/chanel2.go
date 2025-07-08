package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			println("通道输入", i)
		}
		close(ch)
		println("已输出完毕")
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
