package main

import "testing"

func TestBoat_Deliver(t *testing.T) {
	boat := Boat{}
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
    boat.Deliver()
   
}
