package main

import (
	"fmt"
	"sync"
)

// 1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 发送协程
func send(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 接受协程
func receive(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接受到：%d\n", v)
	}
}

// 2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func sendHundred(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}
func receiveHundred(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接受到：%d\n", v)
	}
}

func main() {
	var wg sync.WaitGroup

	// 任务1
	ch1 := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		send(ch1)
	}()
	go func() {
		defer wg.Done()
		receive(ch1)
	}()

	// 任务2
	ch2 := make(chan int, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		sendHundred(ch2)
	}()
	go func() {
		defer wg.Done()
		receiveHundred(ch2)
	}()

	wg.Wait()
}
