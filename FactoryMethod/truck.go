package main

import "fmt"

type Truck struct {
}

func (t *Truck) Deliver() {
	fmt.Println("I am a Truck in the streets delivering stuff")
}