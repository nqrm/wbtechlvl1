package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func (h *Human) printInfo() {
	fmt.Printf("name: %s, age: %d\n", h.name, h.age)
}

type Action struct {
	Human
	Action string
}

func main() {
	p1 := Human{"Bob", 23}
	p2 := Human{"Arnold", 52}

	act1 := Action{p1, "Run"}
	act1.printInfo()

	act2 := Action{Human: p2, Action: "Speak"}
	act2.printInfo()

	act3 := Action{Human{"Вова", 27}, "Drink"}
	act3.printInfo()

}
