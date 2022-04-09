package main

import "fmt"

type RoadLogistics struct {
}

func (r *RoadLogistics) PlanDelivery() {
	fmt.Println("RoadLogistics: ready to road!")
}

func (r *RoadLogistics) CreateTransport() ITransport {
	var myTruck ITransport
	myTruck = &Truck{}

	return myTruck
}