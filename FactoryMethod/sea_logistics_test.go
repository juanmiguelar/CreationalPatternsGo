package main

import "testing"

func Test_SeaLogistics_PlanDelivery(t *testing.T) {
	mySeaLogistics := SeaLogistics{}
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
    mySeaLogistics.PlanDelivery()
   
}


func Test_SeaLogistics_CreateTransport(t *testing.T) {
	mySeaLogistics := SeaLogistics{}
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
    result := mySeaLogistics.CreateTransport()
	if result == nil {
		t.Error("result shoud not be nil")
	}
   
}