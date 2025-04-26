package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {

	fmt.Println(reverseString("snow dog sun", " "))

}

func reverseString(sourceString string, separator string) string {
	strSlice := strings.Split(sourceString, separator)
	reverseStr := ""
	for i := len(strSlice) - 1; i > -1; i-- {
		reverseStr += strSlice[i]
		reverseStr += separator
	}
	return reverseStr
}
