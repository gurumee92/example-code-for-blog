package main

import (
	"fmt"
	"sync"
)

func WriteMapMutex(m map[int]int, num int) {
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
}

func WriteMapChan(m map[int]int, num int) {
	c := make(chan int)

	for i := 0; i < num; i++ {
		go func(i int) {
			c <- i
		}(i)
	}

	for i := 0; i < num; i++ {
		v := <-c
		m[v] = v
	}
}

func main() {
	fmt.Println("map writes start!")
	m := map[int]int{}
	WriteMapMutex(m, 100)
	fmt.Println("map writes end!")
}
