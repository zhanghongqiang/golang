package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go oddNumber()

	wg.Add(1)
	go evenNumber()

	wg.Wait()
}

func oddNumber() {
	defer wg.Done()

	for i := 1; i < 10; i += 2 {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func evenNumber() {
	defer wg.Done()
	for i := 2; i < 10; i += 2 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func runTask(task []Task) {
	for id, t := range task {
		wg.Add(1)
		go func(id int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			elapsed := time.Since(start)
			fmt.Printf("任务%d 执行时间: %v\n", id, elapsed)

		}(id, t)
	}
	wg.Wait()
}
