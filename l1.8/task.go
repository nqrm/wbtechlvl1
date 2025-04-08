package main

import (
	"fmt"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {

	var x int64 = 113
	var bit int64 = 1 // отчет от 0
	res := x ^ (1 << bit)
	fmt.Printf("x =\t%064b = %d\n", x, x)
	fmt.Printf("x1 =\t%064b = %d\n", res, res)

}
