package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	st := make(map[int]string, 10)
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			st[n] = fmt.Sprintf("worker%d", n)
			fmt.Println(n, "-", st[n])
		}(i)
	}

	wg.Wait()
}
