package main

import "fmt"

type Boat struct {
}

func (b *Boat) Deliver() {
	fmt.Println("I am a Boat in the water delivering stuff")
}