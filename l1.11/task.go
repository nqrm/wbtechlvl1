package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	setA := []int{2, 3, 5, 13, 7, 0}
	setB := []int{6, 7, 1, 212, 2, 19, 5}
	var intersection []int

	tempHash := make(map[int]struct{})
	for _, x := range setA {
		tempHash[x] = struct{}{}
	}

	for _, y := range setB {
		if _, ok := tempHash[y]; ok {
			intersection = append(intersection, y)
		}
	}
	fmt.Println(intersection)

	/*
		O(n^2)
		for _, x := range setA {
			for _, y := range setB {
				if x == y {
					intersection = append(intersection, x)
				}
			}
		}*/
}
