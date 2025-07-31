package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func printNumber(oddOrEvenFlag bool, name string) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		if oddOrEvenFlag && i%2 == 0 {
			fmt.Printf("协程：%s，打印输出：%d\n", name, i)
		} else if !oddOrEvenFlag && i%2 == 1 {
			fmt.Printf("协程：%s，打印输出：%d\n", name, i)
		}
		time.Sleep(time.Millisecond * 10)
	}
}

// 2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func taskGroup(name string) {
	defer wg.Done()
	startTime := time.Now()
	sum := 1
	for i := 1; i < 100; i++ {
		sum *= i
	}
	fmt.Printf("协程：%s 完成任务，共耗时：%d\n", name, time.Since(startTime))
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	// // 偶数
	// go printNumber(true, "偶数")
	// // 奇数
	// go printNumber(false, "奇数")

	go taskGroup("任务1")
	go taskGroup("任务2")
	go taskGroup("任务3")

	wg.Wait()
}
