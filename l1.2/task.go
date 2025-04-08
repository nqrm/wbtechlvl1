package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива
(2,4,6,8,10) и выведет их квадраты в stdout.
*/

func worker(wg *sync.WaitGroup, in chan int) {
	defer wg.Done()
	for number := range in {
		fmt.Println(number * number)
	}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	in := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)
	go worker(&wg, in)
	go worker(&wg, in)
	go worker(&wg, in)
	for _, number := range numbers {
		in <- number
	}
	close(in)
	wg.Wait()
}
