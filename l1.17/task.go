package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := binarySearch(arr, 1)
	fmt.Println(index)
}

// поиск до первого вхождения
func binarySearch(arr []int, x int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] < x {
			l = mid + 1
		}
		if arr[mid] > x {
			r = mid - 1
		}
	}

	return -1
}
