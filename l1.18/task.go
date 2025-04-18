package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Increment struct {
	i  int
	mu sync.Mutex
}

func NewIncrement() *Increment {
	return &Increment{}
}

func (inc *Increment) Grow() {
	inc.mu.Lock()
	defer inc.mu.Unlock()
	inc.i += 1
}

func (inc *Increment) Get() int {
	inc.mu.Lock()
	defer inc.mu.Unlock()
	return inc.i
}

func main() {
	i := NewIncrement()
	var wg sync.WaitGroup
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i.Grow()
		}()
	}
	wg.Wait()
	fmt.Println("Итоговое значение счетчика:", i.Get())
}
