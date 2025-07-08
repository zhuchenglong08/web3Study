package main

import "sync"

type Counter struct {
	count int
	mu    *sync.Mutex
}

func (c *Counter) add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	println(c.count)
}
func main() {
	var wg sync.WaitGroup
	muu := sync.Mutex{}
	CC := Counter{1, &muu}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				CC.mu.Lock()
				if CC.count > 100 {
					CC.mu.Unlock()
					break
				}
				println("协程%s", id, "开始执行", CC.count)
				CC.count++
				CC.mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

}
