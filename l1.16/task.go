package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quicksort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	x := arr[start+(end-start)/2]

	l, r := start, end
	for l < r {
		for arr[l] < x {
			l += 1
		}
		for arr[r] > x {
			r -= 1
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l += 1
			r -= 1
		}
	}
	quicksort(arr, start, r)
	quicksort(arr, l, end)

}
