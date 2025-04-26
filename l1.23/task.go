package main

import (
	"errors"
	"fmt"
)

// Удалить i-ый элемент из слайса.

func main() {
	a := []int{2, 13, 5, 10, 8, 11, 44, 76, 21}
	fmt.Println(a)

	b := a[:6]
	fmt.Println(b)
	_ = remove(&b, 2) // a[2] = 5
	fmt.Println(b)
	fmt.Println(a, len(a), cap(a))
	a = a[:cap(a)]
	fmt.Println(a)
}

/*
func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}*/

/*
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}*/

func remove(arr *[]int, index int) error {
	if (0 > index) || (index > len(*arr)-1) {
		return errors.New("index out of range")
	}

	capacity := cap(*arr)
	fullSlice := (*arr)[:capacity]
	for i := index; i < capacity-1; i++ {
		fullSlice[i] = fullSlice[i+1]
	}
	fullSlice[capacity-1] = 0
	*arr = fullSlice[:len(*arr)-1]
	return nil
}
