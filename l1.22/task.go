package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a и b, значение которых > 2^20.
*/

func main() {

	a := big.NewInt(1 << 40)
	b := big.NewInt(1 << 30)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	var result = new(big.Int)
	fmt.Println("a + b =", result.Add(a, b))
	fmt.Println("a * b =", result.Mul(a, b))
	fmt.Println("a / b =", result.Div(a, b))

}
