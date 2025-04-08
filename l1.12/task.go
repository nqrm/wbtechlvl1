package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	stringList := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, x := range stringList {
		set[x] = struct{}{}
	}

	for s := range set {
		fmt.Println(s)
	}

}
