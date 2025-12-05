package main

import (
	"fmt"
	"time"
)

/*
*题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(5 * time.Second)
}
