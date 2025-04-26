package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

/*
1. объект который работает с XML форматом
2. Сам об
*/

type transport interface {
	Navigating() string
}

type Client struct{}

func (c Client) StartNavigating(t transport) string {
	return t.Navigating()
}

type Boat struct{}

func (b Boat) Navigating() string {
	return "Navigating"
}

type Car struct{}

func (c Car) Drive() string {
	return "Driving"
}

type CarAdapter struct {
	car *Car
}

func (adapter CarAdapter) Navigating() string {
	return adapter.car.Drive()
}

func main() {
	client := &Client{}
	boat := &Boat{}
	fmt.Println(client.StartNavigating(boat))
	car := &Car{}
	//client.StartNavigating(car)
	carAdapter := &CarAdapter{car}
	fmt.Println(client.StartNavigating(carAdapter))

}
