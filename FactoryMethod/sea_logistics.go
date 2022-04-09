package main

import "fmt"

type SeaLogistics struct {
}

func (s *SeaLogistics) PlanDelivery() {
	fmt.Println("SeaLogistics: ready to swim!")
}

func (s *SeaLogistics) CreateTransport() ITransport {
	var myBoat ITransport
	myBoat = &Boat{}

	return myBoat
}