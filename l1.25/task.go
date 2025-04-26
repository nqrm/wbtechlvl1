package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	fmt.Println(time.Now())

	Sleep(2)

	fmt.Println(time.Now())

}

func Sleep(seconds int64) {
	sec := time.Duration(seconds) * time.Second

	t := time.NewTimer(sec)
	<-t.C
}
