package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d - int type\n", v)
	case string:
		fmt.Printf("%s - string type\n", v)
	case bool:
		fmt.Printf("%t - bool type\n", v)
	case chan interface{}:
		fmt.Println("channel of interfaces type")
	}
}

func main() {

	ch := make(chan interface{})
	typeSwitch(123)
	typeSwitch("331")
	typeSwitch(false)
	typeSwitch(ch)

	fmt.Println(reflect.TypeOf(ch))
}
