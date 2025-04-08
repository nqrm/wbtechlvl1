package main

import (
	"fmt"
)

/*
Дана последовательность чисел: 2,4,6,8,10.

Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for number := range in {
			out <- number * number
		}
		close(out)
	}()
	return out
}

func generation(numbers *[]int) <-chan int {
	out := make(chan int)
	go func() {
		for _, number := range *numbers {
			out <- number
		}
		close(out)
	}()

	return out
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0

	ch := generation(&numbers)
	for squareNumber := range square(ch) {
		sum += squareNumber
	}

	fmt.Println(sum)
}
