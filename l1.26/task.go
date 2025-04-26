package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	fmt.Printf("abcd - %t\n", isUnique("abcd"))
	fmt.Printf("abCdefAaf - %t\n", isUnique("abCdefAaf"))
	fmt.Printf("aabcd - %t\n", isUnique("aabcd"))
}

func isUnique(s string) bool {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			return false
		} else {
			m[string(s[i])]++
		}
	}
	return true
}
