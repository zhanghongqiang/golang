package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	channelFunc(ch)

	bufferChannelFunc()
}

func channelFunc(ch chan int) {
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func bufferChannelFunc() {
	buffer := 100
	ch := make(chan int, buffer)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= buffer; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("接收到的数字：", v)
		}
	}()

	wg.Wait()
}
