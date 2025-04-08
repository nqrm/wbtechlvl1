package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(ctx context.Context, in chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker killed")
			return
		case number, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(number)
		}
	}
}

func main() {
	fmt.Print("Enter the number of workers: ")
	var workers int
	_, err := fmt.Scan(&workers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	in := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// горутина обрабатывающая сигнал Interrupt. горутины, испольщующие ctx.Done() будут завершаться вместе с ней
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)
		<-exit
		cancel()
	}()

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			worker(ctx, in)
		}()
	}

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Shutting down")
			return
		case in <- rand.Intn(1000):
		}
	}
}
