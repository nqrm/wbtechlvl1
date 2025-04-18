package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	str := "глав🙂рыба"
	tmp := []rune(str)
	out := make([]rune, len(tmp))
	for i := 0; i < len(tmp); i++ {
		out[i] = tmp[len(tmp)-1-i]
	}

	fmt.Println(string(out))

}
