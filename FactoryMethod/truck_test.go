package main

import "testing"

func TestTruck_Deliver(t *testing.T) {
	myTruck := Truck{}
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
    myTruck.Deliver()
   
}
