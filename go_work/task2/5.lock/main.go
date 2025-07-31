package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

// 2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

func main() {

	// counter := NewCounter()

	// var wg sync.WaitGroup

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		defer wg.Done()
	// 		fmt.Printf("协程 %d 开始工作\n", id)

	// 		for j := 0; j < 1000; j++ {
	// 			counter.Increment() // 安全递增
	// 		}

	// 		fmt.Printf("协程 %d 已完成\n", id)
	// 	}(i)
	// }

	// wg.Wait()

	// fmt.Printf("\n最终结果: %d (预期值 = %d)\n",
	// 	counter.GetValue(), 10*1000)

	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("协程 %d 已完成\n", id)
		}(i)
	}

	// 4. 等待所有协程完成
	wg.Wait()

	// 5. 输出最终结果
	fmt.Printf("\n最终结果: %d (预期值 = %d)\n", counter, 10*1000)

}
