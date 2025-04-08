package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func doWorkWithTimeout(ctx context.Context, timeout time.Duration) {
	timeoutCtx, cancelFunc := context.WithTimeout(ctx, timeout)
	defer cancelFunc()

	ch := make(chan int)
	go func() {
		for {
			number := <-ch
			log.Printf("%d\n", number)
		}
	}()

	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(time.Second)
		}
	}()

	<-timeoutCtx.Done()
	log.Println(timeoutCtx.Err())
}

func main() {
	timeout := 5 * time.Second
	doWorkWithTimeout(context.Background(), timeout)
}
