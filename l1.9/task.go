package main

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func sendNumbers(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, x := range nums {
			out <- x
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()

	return out
}

func main() {
	//a := [5]int{4, 10, 7, 12, 9}
	in := sendNumbers(4, 10, 7, 12, 9)
	out := square(in)

	for x := range out {
		fmt.Println(x)
	}

	/*
		for x := range square(sendNumbers(4, 10, 7, 12, 9)) {
			fmt.Println(x)
		}
	*/
}
