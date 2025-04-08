package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func stopWithSignalChan() {
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				log.Println("Worker1...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	stop <- struct{}{}
	log.Println("Worker1 stopped")
	return
}

func stopWithCloseChan() {
	ch := make(chan int)

	go func() {
		for n := range ch {
			fmt.Println("Worker3...", n)
		}
		return
	}()

	ch <- 5
	ch <- 6
	close(ch)
	fmt.Println("Worker3 stopped")
	return
}

func stopWithContext() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				log.Println("Worker2...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	log.Println("Worker2 stopped")

}

func main() {
	stopWithSignalChan()
	stopWithContext()
	stopWithCloseChan()

}
