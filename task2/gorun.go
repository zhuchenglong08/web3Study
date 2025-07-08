package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//任务1
	gotask1()
	gotask2()
	time.Sleep(1 * time.Second)
	//任务2
	sctask1 := newSctask()

	sctask1.addTask(
		func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("任务1")
		},
		func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("任务2")
		},
		func() {
			time.Sleep(400 * time.Millisecond)
			fmt.Println("任务3")
		},
	)
	start := time.Now()
	sctask1.run()
	fmt.Println("耗时:", time.Now().Sub(start))
	sctask1.printTimes()

}

type Task func()
type sctask struct {
	tasks  []Task
	timers map[int]time.Duration
	mu     sync.Mutex
}

func newSctask() *sctask {
	return &sctask{
		timers: make(map[int]time.Duration),
	}
}
func (sc *sctask) addTask(task ...Task) {
	sc.tasks = append(sc.tasks, task...)
}
func (sc *sctask) run() {
	var wg sync.WaitGroup
	wg.Add(len(sc.tasks))
	for i, task := range sc.tasks {
		go func(index int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			colution := time.Since(start)
			sc.mu.Lock()
			sc.timers[index] = colution
			sc.mu.Unlock()
		}(i, task)
	}
	wg.Wait()
}
func (sc *sctask) printTimes() {
	fmt.Println("执行任务时间统计:")
	sc.mu.Lock()
	defer sc.mu.Unlock()
	for i, timer := range sc.timers {
		fmt.Printf("task %d colution %s\n", i, timer)
	}
}

func gotask1() {
	go func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println("协程1 奇数:", i)
		}
	}()
}
func gotask2() {
	go func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println("协程2 偶数:", i)
		}
	}()
}
