package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func main() {
	tasks := []Task{
		makeTask("1"),
		makeTask("2"),
		makeTask("3"),
	}

	durations := run(tasks, 2)

	for i, d := range durations {
		fmt.Printf("Task %d: %v\n", i+1, d)
	}
}

type Task func()

func makeTask(name string) Task {
	return func() {
		duration := time.Duration(rand.Intn(500)) * time.Millisecond
		fmt.Printf("%s start, sleep=%v\n", name, duration)
		time.Sleep(duration)
		fmt.Printf("%s done\n", name)
	}
}

func run(tasks []Task, concurrency int) []time.Duration {
	n := len(tasks)
	durations := make([]time.Duration, n)
	ch := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	wg.Add(n)

	for i, task := range tasks {
		i, task := i, task
		go func() {
			defer wg.Done()

			ch <- struct{}{}
			start := time.Now()
			task()
			durations[i] = time.Since(start)
			<-ch
		}()
	}

	wg.Wait()
	return durations
}
