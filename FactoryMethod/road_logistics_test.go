package main

import "testing"

func Test_RoadLogistics_PlanDelivery(t *testing.T) {
	myRoadLogistics := RoadLogistics{}
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
    myRoadLogistics.PlanDelivery()
   
}


func Test_RoadLogistics_CreateTransport(t *testing.T) {
	myRoadLogistics := RoadLogistics{}
	defer func() {
        if r := recover(); r != nil {
           t.Error("This func should not panic")
        }
    }()
    result := myRoadLogistics.CreateTransport()
	if result == nil {
		t.Error("result shoud not be nil")
	}
   
}