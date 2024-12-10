package main

import "fmt"

type Transport interface {
	Deliver()
}

type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Delivering by land in a truck")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Delivering by sea in a ship")
}

type TransportCreator interface {
	CreateTransport() Transport
}

type TruckCreator struct{}

func (tc *TruckCreator) CreateTransport() Transport {
	return &Truck{}
}

type ShipCreator struct{}

func (sc *ShipCreator) CreateTransport() Transport {
	return &Ship{}
}

func main() {
	var creator TransportCreator

	creator = &TruckCreator{}
	transport := creator.CreateTransport()
	transport.Deliver()

	creator = &ShipCreator{}
	transport = creator.CreateTransport()
	transport.Deliver()
}
