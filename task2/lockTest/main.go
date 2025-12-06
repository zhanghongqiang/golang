package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var mutex sync.Mutex
	var count int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
			fmt.Println(count)
		}()
	}
	wg.Wait()

	atomicTest()

}

func atomicTest() {

	var count int32

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
			fmt.Println(count)
		}()
	}
	wg.Wait()
}
